package GoSFML2

/*
 #include <SFML/Graphics.h>
*/
import "C"
import "unsafe"

/////////////////////////////////////
///		CONSTS
/////////////////////////////////////

const (
	Blend_Alpha    = iota ///< Pixel = Src * a + Dest * (1 - a)
	Blend_Add             ///< Pixel = Src + Dest
	Blend_Multiply        ///< Pixel = Src * Dest
	Blend_None            ///< No blending
)

/////////////////////////////////////
///		STRUCTS
/////////////////////////////////////

type BlendMode int

type RenderStates struct {
	BlendMode BlendMode
	Transform Transform
	texture   *C.sfTexture
	shader    *C.sfShader
}

/////////////////////////////////////
///		GO <-> C
/////////////////////////////////////

func (this *RenderStates) toCPtr() *C.sfRenderStates {
	return (*C.sfRenderStates)(unsafe.Pointer(this))
}