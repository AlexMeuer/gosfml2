/*
	COMPLETE: YES (4.7.2012)
*/

package GoSFML2

// #include <SFML/Graphics.h> 
import "C"
import "runtime"

/////////////////////////////////////
///		STRUCTS
/////////////////////////////////////

type CircleShape struct {
	cptr    *C.sfCircleShape
	texture *Texture //to prevent the GC from deleting the texture
}

/////////////////////////////////////
///		FUNCS
/////////////////////////////////////

func CreateCircleShape(radius float32) *CircleShape {
	shape := &CircleShape{C.sfCircleShape_create(), nil}
	shape.SetRadius(radius)
	runtime.SetFinalizer(shape, (*CircleShape).Destroy)
	return shape
}

func (this *CircleShape) Copy() *CircleShape {
	shape := &CircleShape{C.sfCircleShape_copy(this.cptr), this.texture}
	runtime.SetFinalizer(shape, (*CircleShape).Destroy)
	return shape
}

func (this *CircleShape) Destroy() {
	C.sfCircleShape_destroy(this.cptr)
	this.cptr = nil
}

func (this *CircleShape) SetPosition(pos Vector2f) {
	C.sfCircleShape_setPosition(this.cptr, pos.toC())
}

func (this *CircleShape) SetScale(scale Vector2f) {
	C.sfCircleShape_setScale(this.cptr, scale.toC())
}

func (this *CircleShape) SetOrigin(orig Vector2f) {
	C.sfCircleShape_setOrigin(this.cptr, orig.toC())
}

func (this *CircleShape) SetRotation(rot float32) {
	C.sfCircleShape_setRotation(this.cptr, C.float(rot))
}

func (this *CircleShape) GetRotation() float32 {
	return float32(C.sfCircleShape_getRotation(this.cptr))
}

func (this *CircleShape) GetPosition() (position Vector2f) {
	position.fromC(C.sfCircleShape_getPosition(this.cptr))
	return
}

func (this *CircleShape) GetScale() (scale Vector2f) {
	scale.fromC(C.sfCircleShape_getScale(this.cptr))
	return
}

func (this *CircleShape) GetOrigin() (origin Vector2f) {
	origin.fromC(C.sfCircleShape_getOrigin(this.cptr))
	return
}

func (this *CircleShape) Move(offset Vector2f) {
	C.sfCircleShape_move(this.cptr, offset.toC())
}

func (this *CircleShape) Scale(factor Vector2f) {
	C.sfCircleShape_scale(this.cptr, factor.toC())
}

func (this *CircleShape) Rotate(angle float32) {
	C.sfCircleShape_rotate(this.cptr, C.float(angle))
}

func (this *CircleShape) SetTexture(texture *Texture, resetRect bool) {
	C.sfCircleShape_setTexture(this.cptr, texture.cptr, goBool2C(resetRect))
	this.texture = texture
}

func (this *CircleShape) SetTextureRect(rect Recti) {
	C.sfCircleShape_setTextureRect(this.cptr, rect.toC())
}

func (this *CircleShape) SetFillColor(color Color) {
	C.sfCircleShape_setFillColor(this.cptr, color.toC())
}

func (this *CircleShape) SetOutlineColor(color Color) {
	C.sfCircleShape_setOutlineColor(this.cptr, color.toC())
}

func (this *CircleShape) SetOutlineThickness(thickness float32) {
	C.sfCircleShape_setOutlineThickness(this.cptr, C.float(thickness))
}

func (this *CircleShape) GetTexture() *Texture {
	return this.texture
}

func (this *CircleShape) GetTransform() (transform Transform) {
	transform.fromC(C.sfCircleShape_getTransform(this.cptr))
	return
}

func (this *CircleShape) GetInverseTransform() (transform Transform) {
	transform.fromC(C.sfCircleShape_getInverseTransform(this.cptr))
	return
}

func (this *CircleShape) GetTextureRect() (rect Recti) {
	rect.fromC(C.sfCircleShape_getTextureRect(this.cptr))
	return
}

func (this *CircleShape) GetFillColor() (color Color) {
	color.fromC(C.sfCircleShape_getFillColor(this.cptr))
	return
}

func (this *CircleShape) GetOutlineColor() (color Color) {
	color.fromC(C.sfCircleShape_getOutlineColor(this.cptr))
	return
}

func (this *CircleShape) GetOutlineThickness() float32 {
	return float32(C.sfCircleShape_getOutlineThickness(this.cptr))
}

func (this *CircleShape) GetPointCount() uint {
	return uint(C.sfCircleShape_getPointCount(this.cptr))
}

func (this *CircleShape) GetPoint(index uint) (point Vector2f) {
	point.fromC(C.sfCircleShape_getPoint(this.cptr, C.uint(index)))
	return
}

func (this *CircleShape) SetRadius(radius float32) {
	C.sfCircleShape_setRadius(this.cptr, C.float(radius))
}

func (this *CircleShape) GetRadius() float32 {
	return float32(C.sfCircleShape_getRadius(this.cptr))
}

func (this *CircleShape) SetPointCount(count uint) {
	C.sfCircleShape_setPointCount(this.cptr, C.uint(count))
}

func (this *CircleShape) GetLocalBounds() (rect Rectf) {
	rect.fromC(C.sfCircleShape_getLocalBounds(this.cptr))
	return
}

func (this *CircleShape) GetGlobalBounds() (rect Rectf) {
	rect.fromC(C.sfCircleShape_getGlobalBounds(this.cptr))
	return
}