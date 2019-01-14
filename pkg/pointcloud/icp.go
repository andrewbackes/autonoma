package pointcloud

import (
	"gonum.org/v1/gonum/mat"
	"math"
)

// ICP is an implementation of iterative closest point.
func ICP(source, target *PointCloud, epsilon float64, interations int) *Transformation {
	transformed := source.Copy()
	dist := math.MaxFloat64
	for i := 0; (dist > epsilon) && (i < interations); i++ {
		// for each source point match the closest target point
		var matched *PointCloud
		matched, dist = closestPoints(transformed, target)
		// get transformation that moves source to matched target
		t := findTransformation(transformed, matched)
		// perform transformation on source to move it closer
		transformed = t.Transform(transformed)
	}
	return &Transformation{}
}

func findTransformation(source, matched *PointCloud) *Transformation {
	// compute centroid of the source set (C_s) and matched set (C_m)
	cs := source.Centroid()
	cm := matched.Centroid()
	// get the cross covariance matrix
	cc := crossCovariance(source, matched, cs, cm)
	rot := rotation(cc)
	return &Transformation{
		Rotation:    rot,
		Translation: translation(cs, cm, rot),
	}
}

func crossCovariance(source, matched *PointCloud, sourceCentroid, matchedCentroid Point) mat.Matrix {
	// make new source and match sets with centroids subtracted
	sp := source.Subtract(sourceCentroid)
	mp := matched.Subtract(matchedCentroid)
	// compute cross-covariance matrix
	smatrix := sp.Matrix()
	mmatrix := mp.Matrix()
	crosscovariance := &mat.Dense{}
	crosscovariance.Mul(smatrix, mmatrix.T())
	return crosscovariance
}

// rotation matrix is is R=VU^T
func rotation(crossCovariance mat.Matrix) mat.Matrix {
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
func translation(sourceCentroid, matchedCentroid Point, rotation mat.Matrix) Point {
	var mult mat.Dense
	mult.Mul(rotation, matchedCentroid.ColMatrix())
	return Subtract(sourceCentroid, matToPoint(&mult))
}

func matToPoint(m mat.Matrix) Point {
	r, _ := m.Dims()
	var p Point
	for i := 0; i < r; i++ {
		p.x[i] = m.At(i, 0)
	}
	return p
}

func closestPoints(source, target *PointCloud) (*PointCloud, float64) {
	if len(target.points) < len(source.points) {
		return source, math.MaxFloat64
	}
	closest := &PointCloud{points: make([]Point, 0, len(source.points))}
	distances := float64(0)
	for _, p := range source.points {
		c, d := target.closest(p)
		distances += d
		closest.Add(c)
	}
	aveDist := distances / float64(len(source.points))
	return closest, aveDist
}

// Distance is the shortest distance from a point to the point cloud.
func (p *PointCloud) closest(to Point) (Point, float64) {
	if len(p.points) == 0 {
		return Point{}, math.MaxFloat64
	}
	min := dist(p.points[0], to)
	closest := p.points[0]
	for i := 1; i < len(p.points); i++ {
		d := dist(p.points[i], to)
		if d < min {
			min = d
			closest = p.points[i]
		}
	}
	return closest, min
}
