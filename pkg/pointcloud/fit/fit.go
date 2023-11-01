// Package fit will place a set of vectors correctly into a pointcloud.
package fit

import (
	"fmt"
	"math"

	"github.com/andrewbackes/autonoma/pkg/point"
	"github.com/andrewbackes/autonoma/pkg/point/transformation"
	"github.com/andrewbackes/autonoma/pkg/pointcloud"
	"github.com/montanaflynn/stats"
	"gonum.org/v1/gonum/mat"
)

func ICP(
	source []point.Point,
	origin point.Point,
	target *pointcloud.PointCloud,
	epsilon float64,
	iterations int,
) ([]point.Point, point.Point, float64) {

	// handle starting condition
	if len(target.Points) == 0 {
		return source, origin, 0
	}
	transformed := point.Copy(source)
	transformedOrigin := origin
	dist := math.MaxFloat64
	for i := 0; (dist > epsilon) && (i < iterations); i++ {
		fmt.Println("Iteration", i)
		matched := closestPoints(transformed, target)
		cleanTransformed, cleanMatched := clean(transformed, matched)
		fmt.Println("Cleaned set size:", len(cleanMatched))
		dist = point.AveDistance(cleanTransformed, cleanMatched)
		fmt.Println("Average Distance", dist)
		trans := nextTransformation(cleanTransformed, cleanMatched)
		transformedOrigin = trans.Apply(transformedOrigin)
		apply(trans, transformed)
	}
	return transformed, transformedOrigin, dist
}

func clean(a, b []point.Point) ([]point.Point, []point.Point) {
	k := 1.0
	dists := make([]float64, len(a))
	for i := range a {
		dists[i] = point.Distance(a[i], b[i])
	}
	median, _ := stats.Median(dists)
	fmt.Println("--> median", median)
	outA := make([]point.Point, 0)
	outB := make([]point.Point, 0)
	for i := range a {
		if dists[i] <= k*median {
			outA = append(outA, a[i])
			outB = append(outB, b[i])
		}
	}
	return outA, outB
}

func apply(trans *transformation.Transformation, to []point.Point) {
	for i := range to {
		to[i] = trans.Apply(to[i])
	}
}

func closestPoints(source []point.Point, target *pointcloud.PointCloud) []point.Point {
	matches := make([]point.Point, len(source))
	for i, v := range source {
		matches[i] = closestPoint(v, target)
	}
	return matches
}

func closestPoint(v point.Point, to *pointcloud.PointCloud) point.Point {
	min := math.MaxFloat64
	closest := point.Point{}
	for w := range to.Points {
		d := point.Distance(v, w)
		if d < min {
			min = d
			closest = w
		}
	}
	return closest
}

func closestUniquePoints(source []point.Point, target *pointcloud.PointCloud) []point.Point {
	if len(target.Points) < len(source) {
		panic("not enough target vectors to match")
	}
	used := map[point.Point]struct{}{}
	matches := make([]point.Point, len(source))
	for i, v := range source {
		matches[i] = closestUnusedPoint(v, target, used)
		used[matches[i]] = struct{}{}
	}
	return matches
}

func closestUnusedPoint(v point.Point, to *pointcloud.PointCloud, used map[point.Point]struct{}) point.Point {
	if len(to.Points) <= len(used) {
		panic("no unused points left")
	}
	min := math.MaxFloat64
	closest := point.Point{}
	for w := range to.Points {
		d := point.Distance(v, w)
		if _, taken := used[w]; !taken && d < min {
			min = d
			closest = w
		}
	}
	fmt.Println("--->", v, "closest to", closest)
	return closest
}

func nextTransformation(source, pairs []point.Point) *transformation.Transformation {
	sourceCentroid := point.Centroid(source)
	pairsCentroid := point.Centroid(pairs)
	fmt.Println("--> source centroid", sourceCentroid)
	fmt.Println("--> pairs centroid", pairsCentroid)
	crossCovarianceMatrix := crossCovarianceOf(source, pairs, sourceCentroid, pairsCentroid)
	rotation := rotationOf(crossCovarianceMatrix)
	return &transformation.Transformation{
		Rotation:    rotation,
		Translation: translationOf(sourceCentroid, pairsCentroid, rotation),
	}
}

func crossCovarianceOf(source, pairs []point.Point, sourceCentroid, pairedCentroid point.Point) mat.Matrix {
	// make new source and match sets with centroids subtracted
	shiftedSource := point.Copy(source)
	point.Shift(shiftedSource, sourceCentroid)

	shiftedPairs := point.Copy(pairs)
	point.Shift(shiftedPairs, pairedCentroid)

	// compute cross-covariance matrix
	sourceMatrix := point.Matrix(shiftedSource)
	pairsMatrix := point.Matrix(shiftedPairs)
	crosscovariance := &mat.Dense{}
	crosscovariance.Mul(sourceMatrix, pairsMatrix.T())
	return crosscovariance
}

// rotation matrix is is R=VU^T
func rotationOf(crossCovariance mat.Matrix) mat.Matrix {
	// perform SVD
	svd := &mat.SVD{}
	svd.Factorize(crossCovariance, mat.SVDFull)
	var u, v mat.Dense
	svd.UTo(&u)
	svd.VTo(&v)
	rotation := &mat.Dense{}
	rotation.Mul(&u, v.T())
	return rotation
}

// translation is C_s - RC_m
func translationOf(sourceCentroid, matchedCentroid point.Point, rotation mat.Matrix) point.Point {
	col := matchedCentroid.Matrix()
	var mult mat.Dense
	mult.Mul(rotation, col)
	v := point.Subtract(sourceCentroid, point.FromMatrix(&mult))
	return v
}
