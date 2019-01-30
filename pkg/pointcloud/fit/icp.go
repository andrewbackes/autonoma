package fit

/*
import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"math"
)

// ICP is an implementation of iterative closest point.
func ICP_old(source, target *PointCloud, epsilon float64, interations int) (*PointCloud, *Transformation, float64) {
	transformation := NewTransformation()
	if len(target.Points) == 0 {
		return source, transformation, 0
	}
	transformed := source.Copy()
	dist := math.MaxFloat64
	for i := 0; (dist > epsilon) && (i < interations); i++ {
		// for each source point match the closest target point
		var matched *PointCloud
		matched, dist = closestPoints(transformed, target)
		fmt.Println("---> source points", source.Uniques())
		fmt.Println("---> matched uniques", matched.Uniques())
		fmt.Println("---> error dist", dist)
		// get transformation that moves source to matched target
		transformation = findTransformation(transformed, matched)
		//printMatrix(t.Rotation)
		// perform transformation on source to move it closer
		transformed = transformation.Transform(transformed)
	}
	return transformed, transformation, dist
}

func findTransformation(source, matched *PointCloud) *Transformation {
	// compute centroid of the source set (C_s) and matched set (C_m)
	cs := source.Centroid()
	fmt.Println("source centroid", cs)
	cm := matched.Centroid()
	fmt.Println("matched centroid", cm)
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
	col := matchedCentroid.ColMatrix()
	var mult mat.Dense
	mult.Mul(rotation, col)
	pt := Subtract(sourceCentroid, matToPoint(&mult))
	return pt
}

func closestPoints(source, target *PointCloud) (*PointCloud, float64) {
	if len(target.Points) < len(source.Points) {
		return source, math.MaxFloat64
	}
	closest := &PointCloud{Points: make([]Point, 0, len(source.Points))}
	available := target.Copy()
	distances := float64(0)
	for _, p := range source.Points {
		c, d := target.closest(p)
		distances += d
		closest.Add(c)

	}
	aveDist := distances / float64(len(source.Points))
	return closest, aveDist
}

// Distance is the shortest distance from a point to the point cloud.
func (p *PointCloud) closest(to Point) (Point, float64) {
	if len(p.Points) == 0 {
		return Point{}, math.MaxFloat64
	}
	min := math.MaxFloat64
	closest := Point{}
	for i := 0; i < len(p.Points); i++ {
		d := dist(p.Points[i], to)
		if d < min {
			min = d
			closest = p.Points[i]
		}
	}
	fmt.Println("--->", to, "to", closest)
	return closest, min
}
*/
