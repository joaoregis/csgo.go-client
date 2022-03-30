package vector

func (vector3 *Vector3) CalcVector3WithOtherVector3(otherVector3 Vector3, operation string) *Vector3 {
	switch operation {
	case "+":
		return &Vector3{X: vector3.X + otherVector3.X, Y: vector3.Y + otherVector3.Y, Z: vector3.Z + otherVector3.Z}
	case "-":
		return &Vector3{X: vector3.X - otherVector3.X, Y: vector3.Y - otherVector3.Y, Z: vector3.Z - otherVector3.Z}
	case "*":
		return &Vector3{X: vector3.X * otherVector3.X, Y: vector3.Y * otherVector3.Y, Z: vector3.Z * otherVector3.Z}
	case "/":
		return &Vector3{X: vector3.X / otherVector3.X, Y: vector3.Y / otherVector3.Y, Z: vector3.Z / otherVector3.Z}
	default:
		return &Vector3{}
	}
}

func (vector3 *Vector3) CalcVector3WithOtherValue(newValue float64, operation string) *Vector3 {
	switch operation {
	case "+":
		return &Vector3{X: vector3.X + newValue, Y: vector3.Y + newValue, Z: vector3.Z + newValue}
	case "-":
		return &Vector3{X: vector3.X - newValue, Y: vector3.Y - newValue, Z: vector3.Z - newValue}
	case "*":
		return &Vector3{X: vector3.X * newValue, Y: vector3.Y * newValue, Z: vector3.Z * newValue}
	case "/":
		return &Vector3{X: vector3.X / newValue, Y: vector3.Y / newValue, Z: vector3.Z / newValue}
	default:
		return &Vector3{}
	}
}
