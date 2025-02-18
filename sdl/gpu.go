/*
Portions of this file are copied from the SDL3 headers. The SDL3 license is
reproduced below:

This software is provided 'as-is', without any express or implied
warranty.  In no event will the authors be held liable for any damages
arising from the use of this software.

Permission is granted to anyone to use this software for any purpose,
including commercial applications, and to alter it and redistribute it
freely, subject to the following restrictions:

1. The origin of this software must not be misrepresented; you must not
   claim that you wrote the original software. If you use this software
   in a product, an acknowledgment in the product documentation would be
   appreciated but is not required.
2. Altered source versions must be plainly marked as such, and must not be
   misrepresented as being the original software.
3. This notice may not be removed or altered from any source distribution.
*/

package sdl

// #cgo pkg-config: sdl3
// #cgo noescape SDL_GPUSupportsShaderFormats
// #cgo nocallback SDL_GPUSupportsShaderFormats
// #cgo noescape SDL_GPUSupportsProperties
// #cgo nocallback SDL_GPUSupportsProperties
// #cgo noescape SDL_CreateGPUDevice
// #cgo nocallback SDL_CreateGPUDevice
// #cgo noescape SDL_CreateGPUDeviceWithProperties
// #cgo nocallback SDL_CreateGPUDeviceWithProperties
// #cgo noescape SDL_DestroyGPUDevice
// #cgo nocallback SDL_DestroyGPUDevice
// #cgo noescape SDL_GetNumGPUDrivers
// #cgo nocallback SDL_GetNumGPUDrivers
// #cgo noescape SDL_GetGPUDriver
// #cgo nocallback SDL_GetGPUDriver
// #cgo noescape SDL_GetGPUDeviceDriver
// #cgo nocallback SDL_GetGPUDeviceDriver
// #cgo noescape SDL_GetGPUShaderFormats
// #cgo nocallback SDL_GetGPUShaderFormats
// #cgo noescape SDL_CreateGPUComputePipeline
// #cgo nocallback SDL_CreateGPUComputePipeline
// #cgo noescape SDL_CreateGPUGraphicsPipeline
// #cgo nocallback SDL_CreateGPUGraphicsPipeline
// #cgo noescape SDL_CreateGPUSampler
// #cgo nocallback SDL_CreateGPUSampler
// #cgo noescape SDL_CreateGPUShader
// #cgo nocallback SDL_CreateGPUShader
// #cgo noescape SDL_CreateGPUTexture
// #cgo nocallback SDL_CreateGPUTexture
// #cgo noescape SDL_CreateGPUBuffer
// #cgo nocallback SDL_CreateGPUBuffer
// #cgo noescape SDL_CreateGPUTransferBuffer
// #cgo nocallback SDL_CreateGPUTransferBuffer
// #cgo noescape SDL_SetGPUBufferName
// #cgo nocallback SDL_SetGPUBufferName
// #cgo noescape SDL_SetGPUTextureName
// #cgo nocallback SDL_SetGPUTextureName
// #cgo noescape SDL_InsertGPUDebugLabel
// #cgo nocallback SDL_InsertGPUDebugLabel
// #cgo noescape SDL_PushGPUDebugGroup
// #cgo nocallback SDL_PushGPUDebugGroup
// #cgo noescape SDL_PopGPUDebugGroup
// #cgo nocallback SDL_PopGPUDebugGroup
// #cgo noescape SDL_ReleaseGPUTexture
// #cgo nocallback SDL_ReleaseGPUTexture
// #cgo noescape SDL_ReleaseGPUSampler
// #cgo nocallback SDL_ReleaseGPUSampler
// #cgo noescape SDL_ReleaseGPUBuffer
// #cgo nocallback SDL_ReleaseGPUBuffer
// #cgo noescape SDL_ReleaseGPUTransferBuffer
// #cgo nocallback SDL_ReleaseGPUTransferBuffer
// #cgo noescape SDL_ReleaseGPUComputePipeline
// #cgo nocallback SDL_ReleaseGPUComputePipeline
// #cgo noescape SDL_ReleaseGPUShader
// #cgo nocallback SDL_ReleaseGPUShader
// #cgo noescape SDL_ReleaseGPUGraphicsPipeline
// #cgo nocallback SDL_ReleaseGPUGraphicsPipeline
// #cgo noescape SDL_AcquireGPUCommandBuffer
// #cgo nocallback SDL_AcquireGPUCommandBuffer
// #cgo noescape SDL_PushGPUVertexUniformData
// #cgo nocallback SDL_PushGPUVertexUniformData
// #cgo noescape SDL_PushGPUFragmentUniformData
// #cgo nocallback SDL_PushGPUFragmentUniformData
// #cgo noescape SDL_PushGPUComputeUniformData
// #cgo nocallback SDL_PushGPUComputeUniformData
// #cgo noescape SDL_BeginGPURenderPass
// #cgo nocallback SDL_BeginGPURenderPass
// #cgo noescape SDL_BindGPUGraphicsPipeline
// #cgo nocallback SDL_BindGPUGraphicsPipeline
// #cgo noescape SDL_SetGPUViewport
// #cgo nocallback SDL_SetGPUViewport
// #cgo noescape SDL_SetGPUScissor
// #cgo nocallback SDL_SetGPUScissor
// #cgo noescape SDL_SetGPUBlendConstants
// #cgo nocallback SDL_SetGPUBlendConstants
// #cgo noescape SDL_SetGPUStencilReference
// #cgo nocallback SDL_SetGPUStencilReference
// #cgo noescape SDL_BindGPUVertexBuffers
// #cgo nocallback SDL_BindGPUVertexBuffers
// #cgo noescape SDL_BindGPUIndexBuffer
// #cgo nocallback SDL_BindGPUIndexBuffer
// #cgo noescape SDL_BindGPUVertexSamplers
// #cgo nocallback SDL_BindGPUVertexSamplers
// #cgo noescape SDL_BindGPUVertexStorageTextures
// #cgo nocallback SDL_BindGPUVertexStorageTextures
// #cgo noescape SDL_BindGPUVertexStorageBuffers
// #cgo nocallback SDL_BindGPUVertexStorageBuffers
// #cgo noescape SDL_BindGPUFragmentSamplers
// #cgo nocallback SDL_BindGPUFragmentSamplers
// #cgo noescape SDL_BindGPUFragmentStorageTextures
// #cgo nocallback SDL_BindGPUFragmentStorageTextures
// #cgo noescape SDL_BindGPUFragmentStorageBuffers
// #cgo nocallback SDL_BindGPUFragmentStorageBuffers
// #cgo noescape SDL_DrawGPUIndexedPrimitives
// #cgo nocallback SDL_DrawGPUIndexedPrimitives
// #cgo noescape SDL_DrawGPUPrimitives
// #cgo nocallback SDL_DrawGPUPrimitives
// #cgo noescape SDL_DrawGPUPrimitivesIndirect
// #cgo nocallback SDL_DrawGPUPrimitivesIndirect
// #cgo noescape SDL_DrawGPUIndexedPrimitivesIndirect
// #cgo nocallback SDL_DrawGPUIndexedPrimitivesIndirect
// #cgo noescape SDL_EndGPURenderPass
// #cgo nocallback SDL_EndGPURenderPass
// #cgo noescape SDL_BeginGPUComputePass
// #cgo nocallback SDL_BeginGPUComputePass
// #cgo noescape SDL_BindGPUComputePipeline
// #cgo nocallback SDL_BindGPUComputePipeline
// #cgo noescape SDL_BindGPUComputeSamplers
// #cgo nocallback SDL_BindGPUComputeSamplers
// #cgo noescape SDL_BindGPUComputeStorageTextures
// #cgo nocallback SDL_BindGPUComputeStorageTextures
// #cgo noescape SDL_BindGPUComputeStorageBuffers
// #cgo nocallback SDL_BindGPUComputeStorageBuffers
// #cgo noescape SDL_DispatchGPUCompute
// #cgo nocallback SDL_DispatchGPUCompute
// #cgo noescape SDL_DispatchGPUComputeIndirect
// #cgo nocallback SDL_DispatchGPUComputeIndirect
// #cgo noescape SDL_EndGPUComputePass
// #cgo nocallback SDL_EndGPUComputePass
// #cgo noescape SDL_MapGPUTransferBuffer
// #cgo nocallback SDL_MapGPUTransferBuffer
// #cgo noescape SDL_UnmapGPUTransferBuffer
// #cgo nocallback SDL_UnmapGPUTransferBuffer
// #cgo noescape SDL_BeginGPUCopyPass
// #cgo nocallback SDL_BeginGPUCopyPass
// #cgo noescape SDL_UploadToGPUTexture
// #cgo nocallback SDL_UploadToGPUTexture
// #cgo noescape SDL_UploadToGPUBuffer
// #cgo nocallback SDL_UploadToGPUBuffer
// #cgo noescape SDL_CopyGPUTextureToTexture
// #cgo nocallback SDL_CopyGPUTextureToTexture
// #cgo noescape SDL_CopyGPUBufferToBuffer
// #cgo nocallback SDL_CopyGPUBufferToBuffer
// #cgo noescape SDL_DownloadFromGPUTexture
// #cgo nocallback SDL_DownloadFromGPUTexture
// #cgo noescape SDL_DownloadFromGPUBuffer
// #cgo nocallback SDL_DownloadFromGPUBuffer
// #cgo noescape SDL_EndGPUCopyPass
// #cgo nocallback SDL_EndGPUCopyPass
// #cgo noescape SDL_GenerateMipmapsForGPUTexture
// #cgo nocallback SDL_GenerateMipmapsForGPUTexture
// #cgo noescape SDL_BlitGPUTexture
// #cgo nocallback SDL_BlitGPUTexture
// #cgo noescape SDL_WindowSupportsGPUSwapchainComposition
// #cgo nocallback SDL_WindowSupportsGPUSwapchainComposition
// #cgo noescape SDL_WindowSupportsGPUPresentMode
// #cgo nocallback SDL_WindowSupportsGPUPresentMode
// #cgo noescape SDL_ClaimWindowForGPUDevice
// #cgo nocallback SDL_ClaimWindowForGPUDevice
// #cgo noescape SDL_ReleaseWindowFromGPUDevice
// #cgo nocallback SDL_ReleaseWindowFromGPUDevice
// #cgo noescape SDL_SetGPUSwapchainParameters
// #cgo nocallback SDL_SetGPUSwapchainParameters
// #cgo noescape SDL_SetGPUAllowedFramesInFlight
// #cgo nocallback SDL_SetGPUAllowedFramesInFlight
// #cgo noescape SDL_GetGPUSwapchainTextureFormat
// #cgo nocallback SDL_GetGPUSwapchainTextureFormat
// #cgo noescape SDL_AcquireGPUSwapchainTexture
// #cgo nocallback SDL_AcquireGPUSwapchainTexture
// #cgo noescape SDL_WaitForGPUSwapchain
// #cgo nocallback SDL_WaitForGPUSwapchain
// #cgo noescape SDL_WaitAndAcquireGPUSwapchainTexture
// #cgo nocallback SDL_WaitAndAcquireGPUSwapchainTexture
// #cgo noescape SDL_SubmitGPUCommandBuffer
// #cgo nocallback SDL_SubmitGPUCommandBuffer
// #cgo noescape SDL_SubmitGPUCommandBufferAndAcquireFence
// #cgo nocallback SDL_SubmitGPUCommandBufferAndAcquireFence
// #cgo noescape SDL_CancelGPUCommandBuffer
// #cgo nocallback SDL_CancelGPUCommandBuffer
// #cgo noescape SDL_WaitForGPUIdle
// #cgo nocallback SDL_WaitForGPUIdle
// #cgo noescape SDL_WaitForGPUFences
// #cgo nocallback SDL_WaitForGPUFences
// #cgo noescape SDL_QueryGPUFence
// #cgo nocallback SDL_QueryGPUFence
// #cgo noescape SDL_ReleaseGPUFence
// #cgo nocallback SDL_ReleaseGPUFence
// #cgo noescape SDL_GPUTextureFormatTexelBlockSize
// #cgo nocallback SDL_GPUTextureFormatTexelBlockSize
// #cgo noescape SDL_GPUTextureSupportsFormat
// #cgo nocallback SDL_GPUTextureSupportsFormat
// #cgo noescape SDL_GPUTextureSupportsSampleCount
// #cgo nocallback SDL_GPUTextureSupportsSampleCount
// #cgo noescape SDL_CalculateGPUTextureFormatSize
// #cgo nocallback SDL_CalculateGPUTextureFormatSize
// #include <SDL3/SDL.h>
import "C"
import (
	"runtime"
	"unsafe"
)

// # CategoryGPU
//
// The GPU API offers a cross-platform way for apps to talk to modern graphics
// hardware. It offers both 3D graphics and compute support, in the style of
// Metal, Vulkan, and Direct3D 12.
//
// A basic workflow might be something like this:
//
// The app creates a GPU device with SDL_CreateGPUDevice(), and assigns it to
// a window with SDL_ClaimWindowForGPUDevice()--although strictly speaking you
// can render offscreen entirely, perhaps for image processing, and not use a
// window at all.
//
// Next the app prepares static data (things that are created once and used
// over and over). For example:
//
// - Shaders (programs that run on the GPU): use SDL_CreateGPUShader().
// - Vertex buffers (arrays of geometry data) and other data rendering will
// need: use SDL_UploadToGPUBuffer().
// - Textures (images): use SDL_UploadToGPUTexture().
// - Samplers (how textures should be read from): use SDL_CreateGPUSampler().
// - Render pipelines (precalculated rendering state): use
// SDL_CreateGPUGraphicsPipeline()
//
// To render, the app creates one or more command buffers, with
// SDL_AcquireGPUCommandBuffer(). Command buffers collect rendering
// instructions that will be submitted to the GPU in batch. Complex scenes can
// use multiple command buffers, maybe configured across multiple threads in
// parallel, as long as they are submitted in the correct order, but many apps
// will just need one command buffer per frame.
//
// Rendering can happen to a texture (what other APIs call a "render target")
// or it can happen to the swapchain texture (which is just a special texture
// that represents a window's contents). The app can use
// SDL_WaitAndAcquireGPUSwapchainTexture() to render to the window.
//
// Rendering actually happens in a Render Pass, which is encoded into a
// command buffer. One can encode multiple render passes (or alternate between
// render and compute passes) in a single command buffer, but many apps might
// simply need a single render pass in a single command buffer. Render Passes
// can render to up to four color textures and one depth texture
// simultaneously. If the set of textures being rendered to needs to change,
// the Render Pass must be ended and a new one must be begun.
//
// The app calls SDL_BeginGPURenderPass(). Then it sets states it needs for
// each draw:
//
// - SDL_BindGPUGraphicsPipeline()
// - SDL_SetGPUViewport()
// - SDL_BindGPUVertexBuffers()
// - SDL_BindGPUVertexSamplers()
// - etc
//
// Then, make the actual draw commands with these states:
//
// - SDL_DrawGPUPrimitives()
// - SDL_DrawGPUPrimitivesIndirect()
// - SDL_DrawGPUIndexedPrimitivesIndirect()
// - etc
//
// After all the drawing commands for a pass are complete, the app should call
// SDL_EndGPURenderPass(). Once a render pass ends all render-related state is
// reset.
//
// The app can begin new Render Passes and make new draws in the same command
// buffer until the entire scene is rendered.
//
// Once all of the render commands for the scene are complete, the app calls
// SDL_SubmitGPUCommandBuffer() to send it to the GPU for processing.
//
// If the app needs to read back data from texture or buffers, the API has an
// efficient way of doing this, provided that the app is willing to tolerate
// some latency. When the app uses SDL_DownloadFromGPUTexture() or
// SDL_DownloadFromGPUBuffer(), submitting the command buffer with
// SDL_SubmitGPUCommandBufferAndAcquireFence() will return a fence handle that
// the app can poll or wait on in a thread. Once the fence indicates that the
// command buffer is done processing, it is safe to read the downloaded data.
// Make sure to call SDL_ReleaseGPUFence() when done with the fence.
//
// The API also has "compute" support. The app calls SDL_BeginGPUComputePass()
// with compute-writeable textures and/or buffers, which can be written to in
// a compute shader. Then it sets states it needs for the compute dispatches:
//
// - SDL_BindGPUComputePipeline()
// - SDL_BindGPUComputeStorageBuffers()
// - SDL_BindGPUComputeStorageTextures()
//
// Then, dispatch compute work:
//
// - SDL_DispatchGPUCompute()
//
// For advanced users, this opens up powerful GPU-driven workflows.
//
// Graphics and compute pipelines require the use of shaders, which as
// mentioned above are small programs executed on the GPU. Each backend
// (Vulkan, Metal, D3D12) requires a different shader format. When the app
// creates the GPU device, the app lets the device know which shader formats
// the app can provide. It will then select the appropriate backend depending
// on the available shader formats and the backends available on the platform.
// When creating shaders, the app must provide the correct shader format for
// the selected backend. If you would like to learn more about why the API
// works this way, there is a detailed
// [blog post](https://moonside.games/posts/layers-all-the-way-down/)
// explaining this situation.
//
// It is optimal for apps to pre-compile the shader formats they might use,
// but for ease of use SDL provides a separate project,
// [SDL_shadercross](https://github.com/libsdl-org/SDL_shadercross)
// , for performing runtime shader cross-compilation.
//
// This is an extremely quick overview that leaves out several important
// details. Already, though, one can see that GPU programming can be quite
// complex! If you just need simple 2D graphics, the
// [Render API](https://wiki.libsdl.org/SDL3/CategoryRender)
// is much easier to use but still hardware-accelerated. That said, even for
// 2D applications the performance benefits and expressiveness of the GPU API
// are significant.
//
// The GPU API targets a feature set with a wide range of hardware support and
// ease of portability. It is designed so that the app won't have to branch
// itself by querying feature support. If you need cutting-edge features with
// limited hardware support, this API is probably not for you.
//
// Examples demonstrating proper usage of this API can be found
// [here](https://github.com/TheSpydog/SDL_gpu_examples)
// .
//
// ## Performance considerations
//
// Here are some basic tips for maximizing your rendering performance.
//
// - Beginning a new render pass is relatively expensive. Use as few render
// passes as you can.
// - Minimize the amount of state changes. For example, binding a pipeline is
// relatively cheap, but doing it hundreds of times when you don't need to
// will slow the performance significantly.
// - Perform your data uploads as early as possible in the frame.
// - Don't churn resources. Creating and releasing resources is expensive.
// It's better to create what you need up front and cache it.
// - Don't use uniform buffers for large amounts of data (more than a matrix
// or so). Use a storage buffer instead.
// - Use cycling correctly. There is a detailed explanation of cycling further
// below.
// - Use culling techniques to minimize pixel writes. The less writing the GPU
// has to do the better. Culling can be a very advanced topic but even
// simple culling techniques can boost performance significantly.
//
// In general try to remember the golden rule of performance: doing things is
// more expensive than not doing things. Don't Touch The Driver!
//
// ## FAQ
//
// **Question: When are you adding more advanced features, like ray tracing or
// mesh shaders?**
//
// Answer: We don't have immediate plans to add more bleeding-edge features,
// but we certainly might in the future, when these features prove worthwhile,
// and reasonable to implement across several platforms and underlying APIs.
// So while these things are not in the "never" category, they are definitely
// not "near future" items either.
//
// **Question: Why is my shader not working?**
//
// Answer: A common oversight when using shaders is not properly laying out
// the shader resources/registers correctly. The GPU API is very strict with
// how it wants resources to be laid out and it's difficult for the API to
// automatically validate shaders to see if they have a compatible layout. See
// the documentation for SDL_CreateGPUShader() and
// SDL_CreateGPUComputePipeline() for information on the expected layout.
//
// Another common issue is not setting the correct number of samplers,
// textures, and buffers in SDL_GPUShaderCreateInfo. If possible use shader
// reflection to extract the required information from the shader
// automatically instead of manually filling in the struct's values.
//
// **Question: My application isn't performing very well. Is this the GPU
// API's fault?**
//
// Answer: No. Long answer: The GPU API is a relatively thin layer over the
// underlying graphics API. While it's possible that we have done something
// inefficiently, it's very unlikely especially if you are relatively
// inexperienced with GPU rendering. Please see the performance tips above and
// make sure you are following them. Additionally, tools like RenderDoc can be
// very helpful for diagnosing incorrect behavior and performance issues.
//
// ## System Requirements
//
// **Vulkan:** Supported on Windows, Linux, Nintendo Switch, and certain
// Android devices. Requires Vulkan 1.0 with the following extensions and
// device features:
//
// - `VK_KHR_swapchain`
// - `VK_KHR_maintenance1`
// - `independentBlend`
// - `imageCubeArray`
// - `depthClamp`
// - `shaderClipDistance`
// - `drawIndirectFirstInstance`
//
// **D3D12:** Supported on Windows 10 or newer, Xbox One (GDK), and Xbox
// Series X|S (GDK). Requires a GPU that supports DirectX 12 Feature Level
// 11_1.
//
// **Metal:** Supported on macOS 10.14+ and iOS/tvOS 13.0+. Hardware
// requirements vary by operating system:
//
// - macOS requires an Apple Silicon or
// [Intel Mac2 family](https://developer.apple.com/documentation/metal/mtlfeatureset/mtlfeatureset_macos_gpufamily2_v1?language=objc)
// GPU
// - iOS/tvOS requires an A9 GPU or newer
// - iOS Simulator and tvOS Simulator are unsupported
//
// ## Uniform Data
//
// Uniforms are for passing data to shaders. The uniform data will be constant
// across all executions of the shader.
//
// There are 4 available uniform slots per shader stage (where the stages are
// vertex, fragment, and compute). Uniform data pushed to a slot on a stage
// keeps its value throughout the command buffer until you call the relevant
// Push function on that slot again.
//
// For example, you could write your vertex shaders to read a camera matrix
// from uniform binding slot 0, push the camera matrix at the start of the
// command buffer, and that data will be used for every subsequent draw call.
//
// It is valid to push uniform data during a render or compute pass.
//
// Uniforms are best for pushing small amounts of data. If you are pushing
// more than a matrix or two per call you should consider using a storage
// buffer instead.
//
// ## A Note On Cycling
//
// When using a command buffer, operations do not occur immediately - they
// occur some time after the command buffer is submitted.
//
// When a resource is used in a pending or active command buffer, it is
// considered to be "bound". When a resource is no longer used in any pending
// or active command buffers, it is considered to be "unbound".
//
// If data resources are bound, it is unspecified when that data will be
// unbound unless you acquire a fence when submitting the command buffer and
// wait on it. However, this doesn't mean you need to track resource usage
// manually.
//
// All of the functions and structs that involve writing to a resource have a
// "cycle" bool. SDL_GPUTransferBuffer, SDL_GPUBuffer, and SDL_GPUTexture all
// effectively function as ring buffers on internal resources. When cycle is
// true, if the resource is bound, the cycle rotates to the next unbound
// internal resource, or if none are available, a new one is created. This
// means you don't have to worry about complex state tracking and
// synchronization as long as cycling is correctly employed.
//
// For example: you can call SDL_MapGPUTransferBuffer(), write texture data,
// SDL_UnmapGPUTransferBuffer(), and then SDL_UploadToGPUTexture(). The next
// time you write texture data to the transfer buffer, if you set the cycle
// param to true, you don't have to worry about overwriting any data that is
// not yet uploaded.
//
// Another example: If you are using a texture in a render pass every frame,
// this can cause a data dependency between frames. If you set cycle to true
// in the SDL_GPUColorTargetInfo struct, you can prevent this data dependency.
//
// Cycling will never undefine already bound data. When cycling, all data in
// the resource is considered to be undefined for subsequent commands until
// that data is written again. You must take care not to read undefined data.
//
// Note that when cycling a texture, the entire texture will be cycled, even
// if only part of the texture is used in the call, so you must consider the
// entire texture to contain undefined data after cycling.
//
// You must also take care not to overwrite a section of data that has been
// referenced in a command without cycling first. It is OK to overwrite
// unreferenced data in a bound resource without cycling, but overwriting a
// section of data that has already been referenced will produce unexpected
// results.

// An opaque handle representing the SDL_GPU context.
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GPUDevice
type GPUDevice C.struct_SDL_GPUDevice

// An opaque handle representing a buffer.
//
// Used for vertices, indices, indirect draw commands, and general compute
// data.
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GPUBuffer
type GPUBuffer C.struct_SDL_GPUBuffer

// An opaque handle representing a transfer buffer.
//
// Used for transferring data to and from the device.
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GPUTransferBuffer
type GPUTransferBuffer struct {
	internal *C.struct_SDL_GPUTransferBuffer
	size     int
	mapping  *GPUTransferBufferMapping
}

// An opaque handle representing a texture.
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GPUTexture
type GPUTexture C.struct_SDL_GPUTexture

// An opaque handle representing a sampler.
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GPUSampler
type GPUSampler C.struct_SDL_GPUSampler

// An opaque handle representing a compiled shader object.
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GPUShader
type GPUShader C.struct_SDL_GPUShader

// An opaque handle representing a compute pipeline.
//
// Used during compute passes.
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GPUComputePipeline
type GPUComputePipeline C.struct_SDL_GPUComputePipeline

// An opaque handle representing a graphics pipeline.
//
// Used during render passes.
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GPUGraphicsPipeline
type GPUGraphicsPipeline C.struct_SDL_GPUGraphicsPipeline

// An opaque handle representing a command buffer.
//
// Most state is managed via command buffers. When setting state using a
// command buffer, that state is local to the command buffer.
//
// Commands only begin execution on the GPU once SDL_SubmitGPUCommandBuffer is
// called. Once the command buffer is submitted, it is no longer valid to use
// it.
//
// Command buffers are executed in submission order. If you submit command
// buffer A and then command buffer B all commands in A will begin executing
// before any command in B begins executing.
//
// In multi-threading scenarios, you should only access a command buffer on
// the thread you acquired it from.
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GPUCommandBuffer
type GPUCommandBuffer C.struct_SDL_GPUCommandBuffer

// An opaque handle representing a render pass.
//
// This handle is transient and should not be held or referenced after
// SDL_EndGPURenderPass is called.
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GPURenderPass
type GPURenderPass C.struct_SDL_GPURenderPass

// An opaque handle representing a compute pass.
//
// This handle is transient and should not be held or referenced after
// SDL_EndGPUComputePass is called.
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GPUComputePass
type GPUComputePass C.struct_SDL_GPUComputePass

// An opaque handle representing a copy pass.
//
// This handle is transient and should not be held or referenced after
// SDL_EndGPUCopyPass is called.
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GPUCopyPass
type GPUCopyPass C.struct_SDL_GPUCopyPass

// An opaque handle representing a fence.
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GPUFence
type GPUFence C.struct_SDL_GPUFence

// Specifies the primitive topology of a graphics pipeline.
//
// If you are using POINTLIST you must include a point size output in the
// vertex shader.
//
// - For HLSL compiling to SPIRV you must decorate a float output with
// [[vk::builtin("PointSize")]].
// - For GLSL you must set the gl_PointSize builtin.
// - For MSL you must include a float output with the [[point_size]]
// decorator.
//
// Note that sized point topology is totally unsupported on D3D12. Any size
// other than 1 will be ignored. In general, you should avoid using point
// topology for both compatibility and performance reasons. You WILL regret
// using it.
//
// This enum is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GPUPrimitiveType
type GPUPrimitiveType uint32

const (
	GPUPrimitivetypeTrianglelist  GPUPrimitiveType = iota // A series of separate triangles.
	GPUPrimitivetypeTrianglestrip                         // A series of connected triangles.
	GPUPrimitivetypeLinelist                              // A series of separate lines.
	GPUPrimitivetypeLinestrip                             // A series of connected lines.
	GPUPrimitivetypePointlist                             // A series of separate points.
)

// Specifies how the contents of a texture attached to a render pass are
// treated at the beginning of the render pass.
//
// This enum is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GPULoadOp
type GPULoadOp uint32

const (
	GPULoadopLoad     GPULoadOp = iota // The previous contents of the texture will be preserved.
	GPULoadopClear                     // The contents of the texture will be cleared to a color.
	GPULoadopDontCare                  // The previous contents of the texture need not be preserved. The contents will be undefined.
)

// Specifies how the contents of a texture attached to a render pass are
// treated at the end of the render pass.
//
// This enum is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GPUStoreOp
type GPUStoreOp uint32

const (
	GPUStoreopStore           GPUStoreOp = iota // The contents generated during the render pass will be written to memory.
	GPUStoreopDontCare                          // The contents generated during the render pass are not needed and may be discarded. The contents will be undefined.
	GPUStoreopResolve                           // The multisample contents generated during the render pass will be resolved to a non-multisample texture. The contents in the multisample texture may then be discarded and will be undefined.
	GPUStoreopResolveAndStore                   // The multisample contents generated during the render pass will be resolved to a non-multisample texture. The contents in the multisample texture will be written to memory.
)

// Specifies the size of elements in an index buffer.
//
// This enum is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GPUIndexElementSize
type GPUIndexElementSize uint32

const (
	GPUIndexelementsize16bit GPUIndexElementSize = iota // The index elements are 16-bit.
	GPUIndexelementsize32bit                            // The index elements are 32-bit.
)

// Specifies the pixel format of a texture.
//
// Texture format support varies depending on driver, hardware, and usage
// flags. In general, you should use SDL_GPUTextureSupportsFormat to query if
// a format is supported before using it. However, there are a few guaranteed
// formats.
//
// FIXME: Check universal support for 32-bit component formats FIXME: Check
// universal support for SIMULTANEOUS_READ_WRITE
//
// For SAMPLER usage, the following formats are universally supported:
//
// - R8G8B8A8_UNORM
// - B8G8R8A8_UNORM
// - R8_UNORM
// - R8_SNORM
// - R8G8_UNORM
// - R8G8_SNORM
// - R8G8B8A8_SNORM
// - R16_FLOAT
// - R16G16_FLOAT
// - R16G16B16A16_FLOAT
// - R32_FLOAT
// - R32G32_FLOAT
// - R32G32B32A32_FLOAT
// - R11G11B10_UFLOAT
// - R8G8B8A8_UNORM_SRGB
// - B8G8R8A8_UNORM_SRGB
// - D16_UNORM
//
// For COLOR_TARGET usage, the following formats are universally supported:
//
// - R8G8B8A8_UNORM
// - B8G8R8A8_UNORM
// - R8_UNORM
// - R16_FLOAT
// - R16G16_FLOAT
// - R16G16B16A16_FLOAT
// - R32_FLOAT
// - R32G32_FLOAT
// - R32G32B32A32_FLOAT
// - R8_UINT
// - R8G8_UINT
// - R8G8B8A8_UINT
// - R16_UINT
// - R16G16_UINT
// - R16G16B16A16_UINT
// - R8_INT
// - R8G8_INT
// - R8G8B8A8_INT
// - R16_INT
// - R16G16_INT
// - R16G16B16A16_INT
// - R8G8B8A8_UNORM_SRGB
// - B8G8R8A8_UNORM_SRGB
//
// For STORAGE usages, the following formats are universally supported:
//
// - R8G8B8A8_UNORM
// - R8G8B8A8_SNORM
// - R16G16B16A16_FLOAT
// - R32_FLOAT
// - R32G32_FLOAT
// - R32G32B32A32_FLOAT
// - R8G8B8A8_UINT
// - R16G16B16A16_UINT
// - R8G8B8A8_INT
// - R16G16B16A16_INT
//
// For DEPTH_STENCIL_TARGET usage, the following formats are universally
// supported:
//
// - D16_UNORM
// - Either (but not necessarily both!) D24_UNORM or D32_FLOAT
// - Either (but not necessarily both!) D24_UNORM_S8_UINT or D32_FLOAT_S8_UINT
//
// Unless D16_UNORM is sufficient for your purposes, always check which of
// D24/D32 is supported before creating a depth-stencil texture!
//
// This enum is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GPUTextureFormat
type GPUTextureFormat uint32

const (
	GPUTextureformatInvalid GPUTextureFormat = iota
	GPUTextureformatA8Unorm
	GPUTextureformatR8Unorm
	GPUTextureformatR8G8Unorm
	GPUTextureformatR8G8B8A8Unorm
	GPUTextureformatR16Unorm
	GPUTextureformatR16G16Unorm
	GPUTextureformatR16G16B16A16Unorm
	GPUTextureformatR10G10B10A2Unorm
	GPUTextureformatB5G6R5Unorm
	GPUTextureformatB5G5R5A1Unorm
	GPUTextureformatB4G4R4A4Unorm
	GPUTextureformatB8G8R8A8Unorm
	GPUTextureformatBC1RGBAUnorm
	GPUTextureformatBC2RGBAUnorm
	GPUTextureformatBC3RGBAUnorm
	GPUTextureformatBC4RUnorm
	GPUTextureformatBC5RgUnorm
	GPUTextureformatBC7RGBAUnorm
	GPUTextureformatBC6HRgbFloat
	GPUTextureformatBC6HRgbUfloat
	GPUTextureformatR8Snorm
	GPUTextureformatR8G8Snorm
	GPUTextureformatR8G8B8A8Snorm
	GPUTextureformatR16Snorm
	GPUTextureformatR16G16Snorm
	GPUTextureformatR16G16B16A16Snorm
	GPUTextureformatR16Float
	GPUTextureformatR16G16Float
	GPUTextureformatR16G16B16A16Float
	GPUTextureformatR32Float
	GPUTextureformatR32G32Float
	GPUTextureformatR32G32B32A32Float
	GPUTextureformatR11G11B10Ufloat
	GPUTextureformatR8Uint
	GPUTextureformatR8G8Uint
	GPUTextureformatR8G8B8A8Uint
	GPUTextureformatR16Uint
	GPUTextureformatR16G16Uint
	GPUTextureformatR16G16B16A16Uint
	GPUTextureformatR32Uint
	GPUTextureformatR32G32Uint
	GPUTextureformatR32G32B32A32Uint
	GPUTextureformatR8Int
	GPUTextureformatR8G8Int
	GPUTextureformatR8G8B8A8Int
	GPUTextureformatR16Int
	GPUTextureformatR16G16Int
	GPUTextureformatR16G16B16A16Int
	GPUTextureformatR32Int
	GPUTextureformatR32G32Int
	GPUTextureformatR32G32B32A32Int
	GPUTextureformatR8G8B8A8UnormSrgb
	GPUTextureformatB8G8R8A8UnormSrgb
	GPUTextureformatBC1RGBAUnormSrgb
	GPUTextureformatBC2RGBAUnormSrgb
	GPUTextureformatBC3RGBAUnormSrgb
	GPUTextureformatBC7RGBAUnormSrgb
	GPUTextureformatD16Unorm
	GPUTextureformatD24Unorm
	GPUTextureformatD32Float
	GPUTextureformatD24UnormS8Uint
	GPUTextureformatD32FloatS8Uint
	GPUTextureformatAstc4x4Unorm
	GPUTextureformatAstc5x4Unorm
	GPUTextureformatAstc5x5Unorm
	GPUTextureformatAstc6x5Unorm
	GPUTextureformatAstc6x6Unorm
	GPUTextureformatAstc8x5Unorm
	GPUTextureformatAstc8x6Unorm
	GPUTextureformatAstc8x8Unorm
	GPUTextureformatAstc10x5Unorm
	GPUTextureformatAstc10x6Unorm
	GPUTextureformatAstc10x8Unorm
	GPUTextureformatAstc10x10Unorm
	GPUTextureformatAstc12x10Unorm
	GPUTextureformatAstc12x12Unorm
	GPUTextureformatAstc4x4UnormSrgb
	GPUTextureformatAstc5x4UnormSrgb
	GPUTextureformatAstc5x5UnormSrgb
	GPUTextureformatAstc6x5UnormSrgb
	GPUTextureformatAstc6x6UnormSrgb
	GPUTextureformatAstc8x5UnormSrgb
	GPUTextureformatAstc8x6UnormSrgb
	GPUTextureformatAstc8x8UnormSrgb
	GPUTextureformatAstc10x5UnormSrgb
	GPUTextureformatAstc10x6UnormSrgb
	GPUTextureformatAstc10x8UnormSrgb
	GPUTextureformatAstc10x10UnormSrgb
	GPUTextureformatAstc12x10UnormSrgb
	GPUTextureformatAstc12x12UnormSrgb
	GPUTextureformatAstc4x4Float
	GPUTextureformatAstc5x4Float
	GPUTextureformatAstc5x5Float
	GPUTextureformatAstc6x5Float
	GPUTextureformatAstc6x6Float
	GPUTextureformatAstc8x5Float
	GPUTextureformatAstc8x6Float
	GPUTextureformatAstc8x8Float
	GPUTextureformatAstc10x5Float
	GPUTextureformatAstc10x6Float
	GPUTextureformatAstc10x8Float
	GPUTextureformatAstc10x10Float
	GPUTextureformatAstc12x10Float
	GPUTextureformatAstc12x12Float
)

// Specifies how a texture is intended to be used by the client.
//
// A texture must have at least one usage flag. Note that some usage flag
// combinations are invalid.
//
// With regards to compute storage usage, READ | WRITE means that you can have
// shader A that only writes into the texture and shader B that only reads
// from the texture and bind the same texture to either shader respectively.
// SIMULTANEOUS means that you can do reads and writes within the same shader
// or compute pass. It also implies that atomic ops can be used, since those
// are read-modify-write operations. If you use SIMULTANEOUS, you are
// responsible for avoiding data races, as there is no data synchronization
// within a compute pass. Note that SIMULTANEOUS usage is only supported by a
// limited number of texture formats.
//
// This datatype is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GPUTextureUsageFlags
type GPUTextureUsageFlags uint32

const (
	GPUTextureusageSampler                             GPUTextureUsageFlags = (1 << 0)
	GPUTextureusageColorTarget                         GPUTextureUsageFlags = (1 << 1)
	GPUTextureusageDepthStencilTarget                  GPUTextureUsageFlags = (1 << 2)
	GPUTextureusageGraphicsStorageRead                 GPUTextureUsageFlags = (1 << 3)
	GPUTextureusageComputeStorageRead                  GPUTextureUsageFlags = (1 << 4)
	GPUTextureusageComputeStorageWrite                 GPUTextureUsageFlags = (1 << 5)
	GPUTextureusageComputeStorageSimultaneousReadWrite GPUTextureUsageFlags = (1 << 6)
)

// Specifies the type of a texture.
//
// This enum is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GPUTextureType
type GPUTextureType uint32

const (
	GPUTexturetype2D        GPUTextureType = iota // The texture is a 2-dimensional image.
	GPUTexturetype2DArray                         // The texture is a 2-dimensional array image.
	GPUTexturetype3D                              // The texture is a 3-dimensional image.
	GPUTexturetypeCube                            // The texture is a cube image.
	GPUTexturetypeCubeArray                       // The texture is a cube array image.
)

// Specifies the sample count of a texture.
//
// Used in multisampling. Note that this value only applies when the texture
// is used as a render target.
//
// This enum is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GPUSampleCount
type GPUSampleCount uint32

const (
	GPUSamplecount1 GPUSampleCount = iota // No multisampling.
	GPUSamplecount2                       // MSAA 2x
	GPUSamplecount4                       // MSAA 4x
	GPUSamplecount8                       // MSAA 8x
)

// Specifies the face of a cube map.
//
// Can be passed in as the layer field in texture-related structs.
//
// This enum is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GPUCubeMapFace
type GPUCubeMapFace uint32

const (
	GPUCubemapfacePositivex GPUCubeMapFace = iota
	GPUCubemapfaceNegativex
	GPUCubemapfacePositivey
	GPUCubemapfaceNegativey
	GPUCubemapfacePositivez
	GPUCubemapfaceNegativez
)

// Specifies how a buffer is intended to be used by the client.
//
// A buffer must have at least one usage flag. Note that some usage flag
// combinations are invalid.
//
// Unlike textures, READ | WRITE can be used for simultaneous read-write
// usage. The same data synchronization concerns as textures apply.
//
// This datatype is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GPUBufferUsageFlags
type GPUBufferUsageFlags uint32

const (
	GPUBufferusageVertex              GPUBufferUsageFlags = (1 << 0) /**< Buffer is a vertex buffer. */
	GPUBufferusageIndex               GPUBufferUsageFlags = (1 << 1) /**< Buffer is an index buffer. */
	GPUBufferusageIndirect            GPUBufferUsageFlags = (1 << 2) /**< Buffer is an indirect buffer. */
	GPUBufferusageGraphicsStorageRead GPUBufferUsageFlags = (1 << 3) /**< Buffer supports storage reads in graphics stages. */
	GPUBufferusageComputeStorageRead  GPUBufferUsageFlags = (1 << 4) /**< Buffer supports storage reads in the compute stage. */
	GPUBufferusageComputeStorageWrite GPUBufferUsageFlags = (1 << 5) /**< Buffer supports storage writes in the compute stage. */
)

// Specifies how a transfer buffer is intended to be used by the client.
//
// Note that mapping and copying FROM an upload transfer buffer or TO a
// download transfer buffer is undefined behavior.
//
// This enum is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GPUTransferBufferUsage
type GPUTransferBufferUsage uint32

const (
	GPUTransferbufferusageUpload GPUTransferBufferUsage = iota
	GPUTransferbufferusageDownload
)

// Specifies which stage a shader program corresponds to.
//
// This enum is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GPUShaderStage
type GPUShaderStage uint32

const (
	GPUShaderstageVertex GPUShaderStage = iota
	GPUShaderstageFragment
)

// Specifies the format of shader code.
//
// Each format corresponds to a specific backend that accepts it.
//
// This datatype is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GPUShaderFormat
type GPUShaderFormat uint32

const (
	GPUShaderformatInvalid  GPUShaderFormat = 0
	GPUShaderformatPrivate  GPUShaderFormat = (1 << 0) /**< Shaders for NDA'd platforms. */
	GPUShaderformatSPIRV    GPUShaderFormat = (1 << 1) /**< SPIR-V shaders for Vulkan. */
	GPUShaderformatDXBC     GPUShaderFormat = (1 << 2) /**< DXBC SM5_1 shaders for D3D12. */
	GPUShaderformatDXIL     GPUShaderFormat = (1 << 3) /**< DXIL SM6_0 shaders for D3D12. */
	GPUShaderformatMSL      GPUShaderFormat = (1 << 4) /**< MSL shaders for Metal. */
	GPUShaderformatMetallib GPUShaderFormat = (1 << 5) /**< Precompiled metallib shaders for Metal. */
)

// Specifies the format of a vertex attribute.
//
// This enum is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GPUVertexElementFormat
type GPUVertexElementFormat uint32

const (
	GPUVertexelementformatInvalid GPUVertexElementFormat = iota
	GPUVertexelementformatInt
	GPUVertexelementformatInt2
	GPUVertexelementformatInt3
	GPUVertexelementformatInt4
	GPUVertexelementformatUInt
	GPUVertexelementformatUInt2
	GPUVertexelementformatUInt3
	GPUVertexelementformatUInt4
	GPUVertexelementformatFloat
	GPUVertexelementformatFloat2
	GPUVertexelementformatFloat3
	GPUVertexelementformatFloat4
	GPUVertexelementformatByte2
	GPUVertexelementformatByte4
	GPUVertexelementformatUByte2
	GPUVertexelementformatUByte4
	GPUVertexelementformatByte2Norm
	GPUVertexelementformatByte4Norm
	GPUVertexelementformatUByte2Norm
	GPUVertexelementformatUByte4Norm
	GPUVertexelementformatShort2
	GPUVertexelementformatShort4
	GPUVertexelementformatUShort2
	GPUVertexelementformatUShort4
	GPUVertexelementformatShort2Norm
	GPUVertexelementformatShort4Norm
	GPUVertexelementformatUShort2Norm
	GPUVertexelementformatUShort4Norm
	GPUVertexelementformatHalf2
	GPUVertexelementformatHalf4
)

// Specifies the rate at which vertex attributes are pulled from buffers.
//
// This enum is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GPUVertexInputRate
type GPUVertexInputRate uint32

const (
	GPUVertexinputrateVertex   GPUVertexInputRate = iota // Attribute addressing is a function of the vertex index.
	GPUVertexinputrateInstance                           // Attribute addressing is a function of the instance index.
)

// Specifies the fill mode of the graphics pipeline.
//
// This enum is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GPUFillMode
type GPUFillMode uint32

const (
	GPUFillmodeFill GPUFillMode = iota // Polygons will be rendered via rasterization.
	GPUFillmodeLine                    // Polygon edges will be drawn as line segments.
)

// Specifies the facing direction in which triangle faces will be culled.
//
// This enum is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GPUCullMode
type GPUCullMode uint32

const (
	GPUCullmodeNone  GPUCullMode = iota // No triangles are culled.
	GPUCullmodeFront                    // Front-facing triangles are culled.
	GPUCullmodeBack                     // Back-facing triangles are culled.
)

// Specifies the vertex winding that will cause a triangle to be determined to
// be front-facing.
//
// This enum is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GPUFrontFace
type GPUFrontFace uint32

const (
	GPUFrontfaceCounterClockwise GPUFrontFace = iota // A triangle with counter-clockwise vertex winding will be considered front-facing.
	GPUFrontfaceClockwise                            // A triangle with clockwise vertex winding will be considered front-facing.
)

// Specifies a comparison operator for depth, stencil and sampler operations.
//
// This enum is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GPUCompareOp
type GPUCompareOp uint32

const (
	GPUCompareopInvalid        GPUCompareOp = iota
	GPUCompareopNever                       // The comparison always evaluates false.
	GPUCompareopLess                        // The comparison evaluates reference < test.
	GPUCompareopEqual                       // The comparison evaluates reference == test.
	GPUCompareopLessOrEqual                 // The comparison evaluates reference <= test.
	GPUCompareopGreater                     // The comparison evaluates reference > test.
	GPUCompareopNotEqual                    // The comparison evaluates reference != test.
	GPUCompareopGreaterOrEqual              // The comparison evalutes reference >= test.
	GPUCompareopAlways                      // The comparison always evaluates true.
)

// Specifies what happens to a stored stencil value if stencil tests fail or
// pass.
//
// This enum is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GPUStencilOp
type GPUStencilOp uint32

const (
	GPUStencilopInvalid           GPUStencilOp = iota
	GPUStencilopKeep                           // Keeps the current value.
	GPUStencilopZero                           // Sets the value to 0.
	GPUStencilopReplace                        // Sets the value to reference.
	GPUStencilopIncrementAndClamp              // Increments the current value and clamps to the maximum value.
	GPUStencilopDecrementAndClamp              // Decrements the current value and clamps to 0.
	GPUStencilopInvert                         // Bitwise-inverts the current value.
	GPUStencilopIncrementAndWrap               // Increments the current value and wraps back to 0.
	GPUStencilopDecrementAndWrap               // Decrements the current value and wraps to the maximum value.
)

// Specifies the operator to be used when pixels in a render target are
// blended with existing pixels in the texture.
//
// The source color is the value written by the fragment shader. The
// destination color is the value currently existing in the texture.
//
// This enum is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GPUBlendOp
type GPUBlendOp uint32

const (
	GPUBlendopInvalid         GPUBlendOp = iota
	GPUBlendopAdd                        // (source * source_factor) + (destination * destination_factor)
	GPUBlendopSubtract                   // (source * source_factor) - (destination * destination_factor)
	GPUBlendopReverseSubtract            // (destination * destination_factor) - (source * source_factor)
	GPUBlendopMin                        // min(source, destination)
	GPUBlendopMax                        // max(source, destination)
)

// Specifies a blending factor to be used when pixels in a render target are
// blended with existing pixels in the texture.
//
// The source color is the value written by the fragment shader. The
// destination color is the value currently existing in the texture.
//
// This enum is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GPUBlendFactor
type GPUBlendFactor uint32

const (
	GPUBlendfactorInvalid               GPUBlendFactor = iota
	GPUBlendfactorZero                                 // 0
	GPUBlendfactorOne                                  // 1
	GPUBlendfactorSrcColor                             // source color
	GPUBlendfactorOneMinusSrcColor                     // 1 - source color
	GPUBlendfactorDstColor                             // destination color
	GPUBlendfactorOneMinusDstColor                     // 1 - destination color
	GPUBlendfactorSrcAlpha                             // source alpha
	GPUBlendfactorOneMinusSrcAlpha                     // 1 - source alpha
	GPUBlendfactorDstAlpha                             // destination alpha
	GPUBlendfactorOneMinusDstAlpha                     // 1 - destination alpha
	GPUBlendfactorConstantColor                        // blend constant
	GPUBlendfactorOneMinusConstantColor                // 1 - blend constant
	GPUBlendfactorSrcAlphaSaturate                     // min(source alpha, 1 - destination alpha)
)

// Specifies which color components are written in a graphics pipeline.
//
// This datatype is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GPUColorComponentFlags
type GPUColorComponentFlags byte

const (
	GPUColorcomponentR GPUColorComponentFlags = (1 << 0) /**< the red component */
	GPUColorcomponentG GPUColorComponentFlags = (1 << 1) /**< the green component */
	GPUColorcomponentB GPUColorComponentFlags = (1 << 2) /**< the blue component */
	GPUColorcomponentA GPUColorComponentFlags = (1 << 3) /**< the alpha component */
)

// Specifies a filter operation used by a sampler.
//
// This enum is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GPUFilter
type GPUFilter uint32

const (
	GPUFilterNearest GPUFilter = iota // Point filtering.
	GPUFilterLinear                   // Linear filtering.
)

// Specifies a mipmap mode used by a sampler.
//
// This enum is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GPUSamplerMipmapMode
type GPUSamplerMipmapMode uint32

const (
	GPUSamplermipmapmodeNearest GPUSamplerMipmapMode = iota // Point filtering.
	GPUSamplermipmapmodeLinear                              // Linear filtering.
)

// Specifies behavior of texture sampling when the coordinates exceed the 0-1
// range.
//
// This enum is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GPUSamplerAddressMode
type GPUSamplerAddressMode uint32

const (
	GPUSampleraddressmodeRepeat         GPUSamplerAddressMode = iota // Specifies that the coordinates will wrap around.
	GPUSampleraddressmodeMirroredRepeat                              // Specifies that the coordinates will wrap around mirrored.
	GPUSampleraddressmodeClampToEdge                                 // Specifies that the coordinates will clamp to the 0-1 range.
)

// Specifies the timing that will be used to present swapchain textures to the
// OS.
//
// VSYNC mode will always be supported. IMMEDIATE and MAILBOX modes may not be
// supported on certain systems.
//
// It is recommended to query SDL_WindowSupportsGPUPresentMode after claiming
// the window if you wish to change the present mode to IMMEDIATE or MAILBOX.
//
// - VSYNC: Waits for vblank before presenting. No tearing is possible. If
// there is a pending image to present, the new image is enqueued for
// presentation. Disallows tearing at the cost of visual latency.
// - IMMEDIATE: Immediately presents. Lowest latency option, but tearing may
// occur.
// - MAILBOX: Waits for vblank before presenting. No tearing is possible. If
// there is a pending image to present, the pending image is replaced by the
// new image. Similar to VSYNC, but with reduced visual latency.
//
// This enum is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GPUPresentMode
type GPUPresentMode uint32

const (
	GPUPresentmodeVSync GPUPresentMode = iota
	GPUPresentmodeImmediate
	GPUPresentmodeMailbox
)

// Specifies the texture format and colorspace of the swapchain textures.
//
// SDR will always be supported. Other compositions may not be supported on
// certain systems.
//
// It is recommended to query SDL_WindowSupportsGPUSwapchainComposition after
// claiming the window if you wish to change the swapchain composition from
// SDR.
//
// - SDR: B8G8R8A8 or R8G8B8A8 swapchain. Pixel values are in sRGB encoding.
// - SDR_LINEAR: B8G8R8A8_SRGB or R8G8B8A8_SRGB swapchain. Pixel values are
// stored in memory in sRGB encoding but accessed in shaders in "linear
// sRGB" encoding which is sRGB but with a linear transfer function.
// - HDR_EXTENDED_LINEAR: R16G16B16A16_FLOAT swapchain. Pixel values are in
// extended linear sRGB encoding and permits values outside of the [0, 1]
// range.
// - HDR10_ST2084: A2R10G10B10 or A2B10G10R10 swapchain. Pixel values are in
// BT.2020 ST2084 (PQ) encoding.
//
// This enum is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GPUSwapchainComposition
type GPUSwapchainComposition uint32

const (
	GPUSwapchaincompositionSDR GPUSwapchainComposition = iota
	GPUSwapchaincompositionSDRLinear
	GPUSwapchaincompositionHDRExtendedLinear
	GPUSwapchaincompositionHDR10ST2084
)

// A structure specifying a viewport.
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GPUViewport
type GPUViewport struct {
	X        float32 // The left offset of the viewport.
	Y        float32 // The top offset of the viewport.
	W        float32 // The width of the viewport.
	H        float32 // The height of the viewport.
	MinDepth float32 // The minimum depth of the viewport.
	MaxDepth float32 // The maximum depth of the viewport.
}

// A structure specifying parameters related to transferring data to or from a
// texture.
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GPUTextureTransferInfo
type GPUTextureTransferInfo struct {
	TransferBuffer *GPUTransferBuffer // The transfer buffer used in the transfer operation.
	Offset         int                // The starting byte of the image data in the transfer buffer.
	PixelsPerRow   int                // The number of pixels from one row to the next.
	RowsPerLayer   int                // The number of rows from one layer/depth-slice to the next.
}

// A structure specifying a location in a transfer buffer.
//
// Used when transferring buffer data to or from a transfer buffer.
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GPUTransferBufferLocation
type GPUTransferBufferLocation struct {
	TransferBuffer *GPUTransferBuffer // The transfer buffer used in the transfer operation.
	Offset         int                // The starting byte of the buffer data in the transfer buffer.
}

// A structure specifying a location in a texture.
//
// Used when copying data from one texture to another.
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GPUTextureLocation
type GPUTextureLocation struct {
	Texture  *GPUTexture // The texture used in the copy operation.
	MipLevel int         // The mip level index of the location.
	Layer    int         // The layer index of the location.
	X        int         // The left offset of the location.
	Y        int         // The top offset of the location.
	Z        int         // The front offset of the location.
}

// A structure specifying a region of a texture.
//
// Used when transferring data to or from a texture.
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GPUTextureRegion
type GPUTextureRegion struct {
	Texture  *GPUTexture // The texture used in the copy operation.
	MipLevel int         // The mip level index to transfer.
	Layer    int         // The layer index to transfer.
	X        int         // The left offset of the region.
	Y        int         // The top offset of the region.
	Z        int         // The front offset of the region.
	W        int         // The width of the region.
	H        int         // The height of the region.
	D        int         // The depth of the region.
}

// A structure specifying a region of a texture used in the blit operation.
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GPUBlitRegion
type GPUBlitRegion struct {
	Texture           *GPUTexture // The texture.
	MipLevel          int         // The mip level index of the region.
	LayerOrDepthPlane int         // The layer index or depth plane of the region. This value is treated as a layer index on 2D array and cube textures, and as a depth plane on 3D textures.
	X                 int         // The left offset of the region.
	Y                 int         // The top offset of the region.
	W                 int         // The width of the region.
	H                 int         // The height of the region.
}

// A structure specifying a location in a buffer.
//
// Used when copying data between buffers.
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GPUBufferLocation
type GPUBufferLocation struct {
	Buffer *GPUBuffer // The buffer.
	Offset int        // The starting byte within the buffer.
}

// A structure specifying a region of a buffer.
//
// Used when transferring data to or from buffers.
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GPUBufferRegion
type GPUBufferRegion struct {
	Buffer *GPUBuffer // The buffer.
	Offset int        // The starting byte within the buffer.
	Size   int        // The size in bytes of the region.
}

// A structure specifying the parameters of an indirect draw command.
//
// Note that the `first_vertex` and `first_instance` parameters are NOT
// compatible with built-in vertex/instance ID variables in shaders (for
// example, SV_VertexID); GPU APIs and shader languages do not define these
// built-in variables consistently, so if your shader depends on them, the
// only way to keep behavior consistent and portable is to always pass 0 for
// the correlating parameter in the draw calls.
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GPUIndirectDrawCommand
type GPUIndirectDrawCommand struct {
	NumVertices   uint32 // The number of vertices to draw.
	NumInstances  uint32 // The number of instances to draw.
	FirstVertex   uint32 // The index of the first vertex to draw.
	FirstInstance uint32 // The ID of the first instance to draw.
}

// A structure specifying the parameters of an indexed indirect draw command.
//
// Note that the `first_vertex` and `first_instance` parameters are NOT
// compatible with built-in vertex/instance ID variables in shaders (for
// example, SV_VertexID); GPU APIs and shader languages do not define these
// built-in variables consistently, so if your shader depends on them, the
// only way to keep behavior consistent and portable is to always pass 0 for
// the correlating parameter in the draw calls.
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GPUIndexedIndirectDrawCommand
type GPUIndexedIndirectDrawCommand struct {
	NumIndices    uint32 // The number of indices to draw per instance.
	NumInstances  uint32 // The number of instances to draw.
	FirstIndex    uint32 // The base index within the index buffer.
	VertexOffset  int32  // The value added to the vertex index before indexing into the vertex buffer.
	FirstInstance uint32 // The ID of the first instance to draw.
}

// A structure specifying the parameters of an indexed dispatch command.
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GPUIndirectDispatchCommand
type GPUIndirectDispatchCommand struct {
	GroupcountX uint32 // The number of local workgroups to dispatch in the X dimension.
	GroupcountY uint32 // The number of local workgroups to dispatch in the Y dimension.
	GroupcountZ uint32 // The number of local workgroups to dispatch in the Z dimension.
}

// A structure specifying the parameters of a sampler.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GPUSamplerCreateInfo
type GPUSamplerCreateInfo struct {
	MinFilter        GPUFilter             // The minification filter to apply to lookups.
	MagFilter        GPUFilter             // The magnification filter to apply to lookups.
	MipmapMode       GPUSamplerMipmapMode  // The mipmap filter to apply to lookups.
	AddressModeU     GPUSamplerAddressMode // The addressing mode for U coordinates outside [0, 1).
	AddressModeV     GPUSamplerAddressMode // The addressing mode for V coordinates outside [0, 1).
	AddressModeW     GPUSamplerAddressMode // The addressing mode for W coordinates outside [0, 1).
	MipLodBias       float32               // The bias to be added to mipmap LOD calculation.
	MaxAnisotropy    float32               // The anisotropy value clamp used by the sampler. If enable_anisotropy is false, this is ignored.
	CompareOp        GPUCompareOp          // The comparison operator to apply to fetched data before filtering.
	MinLod           float32               // Clamps the minimum of the computed LOD value.
	MaxLod           float32               // Clamps the maximum of the computed LOD value.
	EnableAnisotropy bool                  // true to enable anisotropic filtering.
	EnableCompare    bool                  // true to enable comparison against a reference value during lookups.
	Props            PropertiesID          // A properties ID for extensions. Should be 0 if no extensions are needed.
}

// A structure specifying the parameters of vertex buffers used in a graphics
// pipeline.
//
// When you call SDL_BindGPUVertexBuffers, you specify the binding slots of
// the vertex buffers. For example if you called SDL_BindGPUVertexBuffers with
// a first_slot of 2 and num_bindings of 3, the binding slots 2, 3, 4 would be
// used by the vertex buffers you pass in.
//
// Vertex attributes are linked to buffers via the buffer_slot field of
// SDL_GPUVertexAttribute. For example, if an attribute has a buffer_slot of
// 0, then that attribute belongs to the vertex buffer bound at slot 0.
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GPUVertexBufferDescription
type GPUVertexBufferDescription struct {
	Slot             int                // The binding slot of the vertex buffer.
	Pitch            int                // The byte pitch between consecutive elements of the vertex buffer.
	InputRate        GPUVertexInputRate // Whether attribute addressing is a function of the vertex index or instance index.
	InstanceStepRate int                // The number of instances to draw using the same per-instance data before advancing in the instance buffer by one element. Ignored unless input_rate is SDL_GPU_VERTEXINPUTRATE_INSTANCE
}

// A structure specifying a vertex attribute.
//
// All vertex attribute locations provided to an SDL_GPUVertexInputState must
// be unique.
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GPUVertexAttribute
type GPUVertexAttribute struct {
	Location   int                    // The shader input location index.
	BufferSlot int                    // The binding slot of the associated vertex buffer.
	Format     GPUVertexElementFormat // The size and type of the attribute data.
	Offset     int                    // The byte offset of this attribute relative to the start of the vertex element.
}

// A structure specifying the parameters of a graphics pipeline vertex input
// state.
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GPUVertexInputState
type GPUVertexInputState struct {
	VertexBufferDescriptions []GPUVertexBufferDescription // A pointer to an array of vertex buffer descriptions.
	VertexAttributes         []GPUVertexAttribute         // A pointer to an array of vertex attribute descriptions.
}

// A structure specifying the stencil operation state of a graphics pipeline.
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GPUStencilOpState
type GPUStencilOpState struct {
	FailOp      GPUStencilOp // The action performed on samples that fail the stencil test.
	PassOp      GPUStencilOp // The action performed on samples that pass the depth and stencil tests.
	DepthFailOp GPUStencilOp // The action performed on samples that pass the stencil test and fail the depth test.
	CompareOp   GPUCompareOp // The comparison operator used in the stencil test.
}

// A structure specifying the blend state of a color target.
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GPUColorTargetBlendState
type GPUColorTargetBlendState struct {
	SrcColorBlendfactor  GPUBlendFactor         // The value to be multiplied by the source RGB value.
	DstColorBlendfactor  GPUBlendFactor         // The value to be multiplied by the destination RGB value.
	ColorBlendOp         GPUBlendOp             // The blend operation for the RGB components.
	SrcAlphaBlendfactor  GPUBlendFactor         // The value to be multiplied by the source alpha.
	DstAlphaBlendfactor  GPUBlendFactor         // The value to be multiplied by the destination alpha.
	AlphaBlendOp         GPUBlendOp             // The blend operation for the alpha component.
	ColorWriteMask       GPUColorComponentFlags // A bitmask specifying which of the RGBA components are enabled for writing. Writes to all channels if enable_color_write_mask is false.
	EnableBlend          bool                   // Whether blending is enabled for the color target.
	EnableColorWriteMask bool                   // Whether the color write mask is enabled.
}

// A structure specifying code and metadata for creating a shader object.
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GPUShaderCreateInfo
type GPUShaderCreateInfo struct {
	Code               []byte          // A pointer to shader code.
	Entrypoint         string          // A pointer to a null-terminated UTF-8 string specifying the entry point function name for the shader.
	Format             GPUShaderFormat // The format of the shader code.
	Stage              GPUShaderStage  // The stage the shader program corresponds to.
	NumSamplers        int             // The number of samplers defined in the shader.
	NumStorageTextures int             // The number of storage textures defined in the shader.
	NumStorageBuffers  int             // The number of storage buffers defined in the shader.
	NumUniformBuffers  int             // The number of uniform buffers defined in the shader.
	Props              PropertiesID    // A properties ID for extensions. Should be 0 if no extensions are needed.
}

// A structure specifying the parameters of a texture.
//
// Usage flags can be bitwise OR'd together for combinations of usages. Note
// that certain usage combinations are invalid, for example SAMPLER and
// GRAPHICS_STORAGE.
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GPUTextureCreateInfo
type GPUTextureCreateInfo struct {
	Type              GPUTextureType       // The base dimensionality of the texture.
	Format            GPUTextureFormat     // The pixel format of the texture.
	Usage             GPUTextureUsageFlags // How the texture is intended to be used by the client.
	Width             int                  // The width of the texture.
	Height            int                  // The height of the texture.
	LayerCountOrDepth int                  // The layer count or depth of the texture. This value is treated as a layer count on 2D array textures, and as a depth value on 3D textures.
	NumLevels         int                  // The number of mip levels in the texture.
	SampleCount       GPUSampleCount       // The number of samples per texel. Only applies if the texture is used as a render target.
	Props             PropertiesID         // A properties ID for extensions. Should be 0 if no extensions are needed.
}

// A structure specifying the parameters of a buffer.
//
// Usage flags can be bitwise OR'd together for combinations of usages. Note
// that certain combinations are invalid, for example VERTEX and INDEX.
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GPUBufferCreateInfo
type GPUBufferCreateInfo struct {
	Usage GPUBufferUsageFlags // How the buffer is intended to be used by the client.
	Size  int                 // The size in bytes of the buffer.
	Props PropertiesID        // A properties ID for extensions. Should be 0 if no extensions are needed.
}

// A structure specifying the parameters of a transfer buffer.
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GPUTransferBufferCreateInfo
type GPUTransferBufferCreateInfo struct {
	Usage GPUTransferBufferUsage // How the transfer buffer is intended to be used by the client.
	Size  int                    // The size in bytes of the transfer buffer.
	Props PropertiesID           // A properties ID for extensions. Should be 0 if no extensions are needed.
}

// A structure specifying the parameters of the graphics pipeline rasterizer
// state.
//
// NOTE: Some backend APIs (D3D11/12) will enable depth clamping even if
// enable_depth_clip is true. If you rely on this clamp+clip behavior,
// consider enabling depth clip and then manually clamping depth in your
// fragment shaders on Metal and Vulkan.
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GPURasterizerState
type GPURasterizerState struct {
	FillMode                GPUFillMode  // Whether polygons will be filled in or drawn as lines.
	CullMode                GPUCullMode  // The facing direction in which triangles will be culled.
	FrontFace               GPUFrontFace // The vertex winding that will cause a triangle to be determined as front-facing.
	DepthBiasConstantFactor float32      // A scalar factor controlling the depth value added to each fragment.
	DepthBiasClamp          float32      // The maximum depth bias of a fragment.
	DepthBiasSlopeFactor    float32      // A scalar factor applied to a fragment's slope in depth calculations.
	EnableDepthBias         bool         // true to bias fragment depth values.
	EnableDepthClip         bool         // true to enable depth clip, false to enable depth clamp.
}

// A structure specifying the parameters of the graphics pipeline multisample
// state.
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GPUMultisampleState
type GPUMultisampleState struct {
	SampleCount GPUSampleCount // The number of samples to be used in rasterization.
	SampleMask  uint32         // Determines which samples get updated in the render targets. Treated as 0xFFFFFFFF if enable_mask is false.
	EnableMask  bool           // Enables sample masking.
}

// A structure specifying the parameters of the graphics pipeline depth
// stencil state.
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GPUDepthStencilState
type GPUDepthStencilState struct {
	CompareOp         GPUCompareOp      // The comparison operator used for depth testing.
	BackStencilState  GPUStencilOpState // The stencil op state for back-facing triangles.
	FrontStencilState GPUStencilOpState // The stencil op state for front-facing triangles.
	CompareMask       byte              // Selects the bits of the stencil values participating in the stencil test.
	WriteMask         byte              // Selects the bits of the stencil values updated by the stencil test.
	EnableDepthTest   bool              // true enables the depth test.
	EnableDepthWrite  bool              // true enables depth writes. Depth writes are always disabled when enable_depth_test is false.
	EnableStencilTest bool              // true enables the stencil test.
}

// A structure specifying the parameters of color targets used in a graphics
// pipeline.
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GPUColorTargetDescription
type GPUColorTargetDescription struct {
	Format     GPUTextureFormat         // The pixel format of the texture to be used as a color target.
	BlendState GPUColorTargetBlendState // The blend state to be used for the color target.
}

// A structure specifying the descriptions of render targets used in a
// graphics pipeline.
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GPUGraphicsPipelineTargetInfo
type GPUGraphicsPipelineTargetInfo struct {
	ColorTargetDescriptions []GPUColorTargetDescription // A pointer to an array of color target descriptions.
	DepthStencilFormat      GPUTextureFormat            // The pixel format of the depth-stencil target. Ignored if has_depth_stencil_target is false.
	HasDepthStencilTarget   bool                        // true specifies that the pipeline uses a depth-stencil target.
}

// A structure specifying the parameters of a graphics pipeline state.
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GPUGraphicsPipelineCreateInfo
type GPUGraphicsPipelineCreateInfo struct {
	VertexShader      *GPUShader                    // The vertex shader used by the graphics pipeline.
	FragmentShader    *GPUShader                    // The fragment shader used by the graphics pipeline.
	VertexInputState  GPUVertexInputState           // The vertex layout of the graphics pipeline.
	PrimitiveType     GPUPrimitiveType              // The primitive topology of the graphics pipeline.
	RasterizerState   GPURasterizerState            // The rasterizer state of the graphics pipeline.
	MultisampleState  GPUMultisampleState           // The multisample state of the graphics pipeline.
	DepthStencilState GPUDepthStencilState          // The depth-stencil state of the graphics pipeline.
	TargetInfo        GPUGraphicsPipelineTargetInfo // Formats and blend modes for the render targets of the graphics pipeline.
	Props             PropertiesID                  // A properties ID for extensions. Should be 0 if no extensions are needed.
}

// A structure specifying the parameters of a compute pipeline state.
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GPUComputePipelineCreateInfo
type GPUComputePipelineCreateInfo struct {
	Code                        []byte          // A pointer to compute shader code.
	Entrypoint                  string          // A pointer to a null-terminated UTF-8 string specifying the entry point function name for the shader.
	Format                      GPUShaderFormat // The format of the compute shader code.
	NumSamplers                 uint            // The number of samplers defined in the shader.
	NumReadonlyStorageTextures  uint            // The number of readonly storage textures defined in the shader.
	NumReadonlyStorageBuffers   uint            // The number of readonly storage buffers defined in the shader.
	NumReadwriteStorageTextures uint            // The number of read-write storage textures defined in the shader.
	NumReadwriteStorageBuffers  uint            // The number of read-write storage buffers defined in the shader.
	NumUniformBuffers           uint            // The number of uniform buffers defined in the shader.
	ThreadcountX                uint            // The number of threads in the X dimension. This should match the value in the shader.
	ThreadcountY                uint            // The number of threads in the Y dimension. This should match the value in the shader.
	ThreadcountZ                uint            // The number of threads in the Z dimension. This should match the value in the shader.
	Props                       PropertiesID    // A properties ID for extensions. Should be 0 if no extensions are needed.
}

// A structure specifying the parameters of a color target used by a render
// pass.
//
// The load_op field determines what is done with the texture at the beginning
// of the render pass.
//
// - LOAD: Loads the data currently in the texture. Not recommended for
// multisample textures as it requires significant memory bandwidth.
// - CLEAR: Clears the texture to a single color.
// - DONT_CARE: The driver will do whatever it wants with the texture memory.
// This is a good option if you know that every single pixel will be touched
// in the render pass.
//
// The store_op field determines what is done with the color results of the
// render pass.
//
// - STORE: Stores the results of the render pass in the texture. Not
// recommended for multisample textures as it requires significant memory
// bandwidth.
// - DONT_CARE: The driver will do whatever it wants with the texture memory.
// This is often a good option for depth/stencil textures.
// - RESOLVE: Resolves a multisample texture into resolve_texture, which must
// have a sample count of 1. Then the driver may discard the multisample
// texture memory. This is the most performant method of resolving a
// multisample target.
// - RESOLVE_AND_STORE: Resolves a multisample texture into the
// resolve_texture, which must have a sample count of 1. Then the driver
// stores the multisample texture's contents. Not recommended as it requires
// significant memory bandwidth.
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GPUColorTargetInfo
type GPUColorTargetInfo struct {
	Texture             *GPUTexture // The texture that will be used as a color target by a render pass.
	MipLevel            int         // The mip level to use as a color target.
	LayerOrDepthPlane   int         // The layer index or depth plane to use as a color target. This value is treated as a layer index on 2D array and cube textures, and as a depth plane on 3D textures.
	ClearColor          FColor      // The color to clear the color target to at the start of the render pass. Ignored if SDL_GPU_LOADOP_CLEAR is not used.
	LoadOp              GPULoadOp   // What is done with the contents of the color target at the beginning of the render pass.
	StoreOp             GPUStoreOp  // What is done with the results of the render pass.
	ResolveTexture      *GPUTexture // The texture that will receive the results of a multisample resolve operation. Ignored if a RESOLVE* store_op is not used.
	ResolveMipLevel     int         // The mip level of the resolve texture to use for the resolve operation. Ignored if a RESOLVE* store_op is not used.
	ResolveLayer        int         // The layer index of the resolve texture to use for the resolve operation. Ignored if a RESOLVE* store_op is not used.
	Cycle               bool        // true cycles the texture if the texture is bound and load_op is not LOAD
	CycleResolveTexture bool        // true cycles the resolve texture if the resolve texture is bound. Ignored if a RESOLVE* store_op is not used.
}

// A structure specifying the parameters of a depth-stencil target used by a
// render pass.
//
// The load_op field determines what is done with the depth contents of the
// texture at the beginning of the render pass.
//
// - LOAD: Loads the depth values currently in the texture.
// - CLEAR: Clears the texture to a single depth.
// - DONT_CARE: The driver will do whatever it wants with the memory. This is
// a good option if you know that every single pixel will be touched in the
// render pass.
//
// The store_op field determines what is done with the depth results of the
// render pass.
//
// - STORE: Stores the depth results in the texture.
// - DONT_CARE: The driver will do whatever it wants with the depth results.
// This is often a good option for depth/stencil textures that don't need to
// be reused again.
//
// The stencil_load_op field determines what is done with the stencil contents
// of the texture at the beginning of the render pass.
//
// - LOAD: Loads the stencil values currently in the texture.
// - CLEAR: Clears the stencil values to a single value.
// - DONT_CARE: The driver will do whatever it wants with the memory. This is
// a good option if you know that every single pixel will be touched in the
// render pass.
//
// The stencil_store_op field determines what is done with the stencil results
// of the render pass.
//
// - STORE: Stores the stencil results in the texture.
// - DONT_CARE: The driver will do whatever it wants with the stencil results.
// This is often a good option for depth/stencil textures that don't need to
// be reused again.
//
// Note that depth/stencil targets do not support multisample resolves.
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GPUDepthStencilTargetInfo
type GPUDepthStencilTargetInfo struct {
	Texture        *GPUTexture // The texture that will be used as the depth stencil target by the render pass.
	ClearDepth     float32     // The value to clear the depth component to at the beginning of the render pass. Ignored if SDL_GPU_LOADOP_CLEAR is not used.
	LoadOp         GPULoadOp   // What is done with the depth contents at the beginning of the render pass.
	StoreOp        GPUStoreOp  // What is done with the depth results of the render pass.
	StencilLoadOp  GPULoadOp   // What is done with the stencil contents at the beginning of the render pass.
	StencilStoreOp GPUStoreOp  // What is done with the stencil results of the render pass.
	Cycle          bool        // true cycles the texture if the texture is bound and any load ops are not LOAD
	ClearStencil   byte        // The value to clear the stencil component to at the beginning of the render pass. Ignored if SDL_GPU_LOADOP_CLEAR is not used.
}

// A structure containing parameters for a blit command.
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GPUBlitInfo
type GPUBlitInfo struct {
	Source      GPUBlitRegion // The source region for the blit.
	Destination GPUBlitRegion // The destination region for the blit.
	LoadOp      GPULoadOp     // What is done with the contents of the destination before the blit.
	ClearColor  FColor        // The color to clear the destination region to before the blit. Ignored if load_op is not SDL_GPU_LOADOP_CLEAR.
	FlipMode    FlipMode      // The flip mode for the source region.
	Filter      GPUFilter     // The filter mode used when blitting.
	Cycle       bool          // true cycles the destination texture if it is already bound.
}

// A structure specifying parameters in a buffer binding call.
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GPUBufferBinding
type GPUBufferBinding struct {
	Buffer *GPUBuffer // The buffer to bind. Must have been created with SDL_GPU_BUFFERUSAGE_VERTEX for SDL_BindGPUVertexBuffers, or SDL_GPU_BUFFERUSAGE_INDEX for SDL_BindGPUIndexBuffer.
	Offset int        // The starting byte of the data to bind in the buffer.
}

// A structure specifying parameters in a sampler binding call.
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GPUTextureSamplerBinding
type GPUTextureSamplerBinding struct {
	Texture *GPUTexture // The texture to bind. Must have been created with SDL_GPU_TEXTUREUSAGE_SAMPLER.
	Sampler *GPUSampler // The sampler to bind.
}

// A structure specifying parameters related to binding buffers in a compute
// pass.
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GPUStorageBufferReadWriteBinding
type GPUStorageBufferReadWriteBinding struct {
	Buffer *GPUBuffer // The buffer to bind. Must have been created with SDL_GPU_BUFFERUSAGE_COMPUTE_STORAGE_WRITE.
	Cycle  bool       // true cycles the buffer if it is already bound.
}

// A structure specifying parameters related to binding textures in a compute
// pass.
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GPUStorageTextureReadWriteBinding
type GPUStorageTextureReadWriteBinding struct {
	Texture  *GPUTexture // The texture to bind. Must have been created with SDL_GPU_TEXTUREUSAGE_COMPUTE_STORAGE_WRITE or SDL_GPU_TEXTUREUSAGE_COMPUTE_STORAGE_SIMULTANEOUS_READ_WRITE.
	MipLevel int         // The mip level index to bind.
	Layer    int         // The layer index to bind.
	Cycle    bool        // true cycles the texture if it is already bound.
}

// Checks for GPU runtime support.
//
// format_flags: a bitflag indicating which shader formats the app is
// able to provide.
//
// name: the preferred GPU driver, or NULL to let SDL pick the optimal
// driver.
//
// Returns true if supported, false otherwise.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GPUSupportsShaderFormats
func GPUSupportsShaderFormats(formatFlags GPUShaderFormat, name string) bool {
	return (bool)(C.SDL_GPUSupportsShaderFormats((C.SDL_GPUShaderFormat)(formatFlags), tmpstring(name)))
}

// Checks for GPU runtime support.
//
// props: the properties to use.
//
// Returns true if supported, false otherwise.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GPUSupportsProperties
func GPUSupportsProperties(props PropertiesID) bool {
	return (bool)(C.SDL_GPUSupportsProperties((C.SDL_PropertiesID)(props)))
}

// Creates a GPU context.
//
// format_flags: a bitflag indicating which shader formats the app is
// able to provide.
//
// debug_mode: enable debug mode properties and validations.
//
// name: the preferred GPU driver, or NULL to let SDL pick the optimal
// driver.
//
// Returns a GPU context on success or NULL on failure; call SDL_GetError()
// for more information.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_CreateGPUDevice
func CreateGPUDevice(formatFlags GPUShaderFormat, debugMode bool, name string) (*GPUDevice, error) {
	dev := (*GPUDevice)(C.SDL_CreateGPUDevice((C.SDL_GPUShaderFormat)(formatFlags), (C.bool)(debugMode), tmpstring(name)))
	if dev == nil {
		return nil, getError()
	}
	return dev, nil
}

// Creates a GPU context.
//
// These are the supported properties:
//
// - `SDL_PROP_GPU_DEVICE_CREATE_DEBUGMODE_BOOLEAN`: enable debug mode
// properties and validations, defaults to true.
// - `SDL_PROP_GPU_DEVICE_CREATE_PREFERLOWPOWER_BOOLEAN`: enable to prefer
// energy efficiency over maximum GPU performance, defaults to false.
// - `SDL_PROP_GPU_DEVICE_CREATE_NAME_STRING`: the name of the GPU driver to
// use, if a specific one is desired.
//
// These are the current shader format properties:
//
// - `SDL_PROP_GPU_DEVICE_CREATE_SHADERS_PRIVATE_BOOLEAN`: The app is able to
// provide shaders for an NDA platform.
// - `SDL_PROP_GPU_DEVICE_CREATE_SHADERS_SPIRV_BOOLEAN`: The app is able to
// provide SPIR-V shaders if applicable.
// - `SDL_PROP_GPU_DEVICE_CREATE_SHADERS_DXBC_BOOLEAN`: The app is able to
// provide DXBC shaders if applicable
// - `SDL_PROP_GPU_DEVICE_CREATE_SHADERS_DXIL_BOOLEAN`: The app is able to
// provide DXIL shaders if applicable.
// - `SDL_PROP_GPU_DEVICE_CREATE_SHADERS_MSL_BOOLEAN`: The app is able to
// provide MSL shaders if applicable.
// - `SDL_PROP_GPU_DEVICE_CREATE_SHADERS_METALLIB_BOOLEAN`: The app is able to
// provide Metal shader libraries if applicable.
//
// With the D3D12 renderer:
//
// - `SDL_PROP_GPU_DEVICE_CREATE_D3D12_SEMANTIC_NAME_STRING`: the prefix to
// use for all vertex semantics, default is "TEXCOORD".
//
// props: the properties to use.
//
// Returns a GPU context on success or NULL on failure; call SDL_GetError()
// for more information.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_CreateGPUDeviceWithProperties
func CreateGPUDeviceWithProperties(props PropertiesID) (*GPUDevice, error) {
	dev := (*GPUDevice)(C.SDL_CreateGPUDeviceWithProperties((C.SDL_PropertiesID)(props)))
	if dev == nil {
		return nil, getError()
	}
	return dev, nil
}

const PropGPUDeviceCreateDebugmodeBoolean = "SDL.gpu.device.create.debugmode"
const PropGPUDeviceCreatePreferlowpowerBoolean = "SDL.gpu.device.create.preferlowpower"
const PropGPUDeviceCreateNameString = "SDL.gpu.device.create.name"
const PropGPUDeviceCreateShadersPrivateBoolean = "SDL.gpu.device.create.shaders.private"
const PropGPUDeviceCreateShadersSpirvBoolean = "SDL.gpu.device.create.shaders.spirv"
const PropGPUDeviceCreateShadersDxbcBoolean = "SDL.gpu.device.create.shaders.dxbc"
const PropGPUDeviceCreateShadersDxilBoolean = "SDL.gpu.device.create.shaders.dxil"
const PropGPUDeviceCreateShadersMslBoolean = "SDL.gpu.device.create.shaders.msl"
const PropGPUDeviceCreateShadersMetallibBoolean = "SDL.gpu.device.create.shaders.metallib"
const PropGPUDeviceCreateD3D12SemanticNameString = "SDL.gpu.device.create.d3d12.semantic"

// Destroys a GPU context previously returned by SDL_CreateGPUDevice.
//
// device: a GPU Context to destroy.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_DestroyGPUDevice
func (device *GPUDevice) Destroy() {
	C.SDL_DestroyGPUDevice((*C.SDL_GPUDevice)(device))
}

// Get the number of GPU drivers compiled into SDL.
//
// Returns the number of built in GPU drivers.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetNumGPUDrivers
func GetNumGPUDrivers() int {
	return (int)(C.SDL_GetNumGPUDrivers())
}

// Get the name of a built in GPU driver.
//
// The GPU drivers are presented in the order in which they are normally
// checked during initialization.
//
// The names of drivers are all simple, low-ASCII identifiers, like "vulkan",
// "metal" or "direct3d12". These never have Unicode characters, and are not
// meant to be proper names.
//
// index: the index of a GPU driver.
//
// Returns the name of the GPU driver with the given **index**.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetGPUDriver
func GetGPUDriver(index int) string {
	return C.GoString(C.SDL_GetGPUDriver((C.int)(index)))
}

// Returns the name of the backend used to create this GPU context.
//
// device: a GPU context to query.
//
// Returns the name of the device's driver, or NULL on error.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetGPUDeviceDriver
func (device *GPUDevice) Driver() string {
	return C.GoString(C.SDL_GetGPUDeviceDriver((*C.SDL_GPUDevice)(device)))
}

// Returns the supported shader formats for this GPU context.
//
// device: a GPU context to query.
//
// Returns a bitflag indicating which shader formats the driver is able to
// consume.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetGPUShaderFormats
func (device *GPUDevice) ShaderFormats() GPUShaderFormat {
	return (GPUShaderFormat)(C.SDL_GetGPUShaderFormats((*C.SDL_GPUDevice)(device)))
}

// Creates a pipeline object to be used in a compute workflow.
//
// Shader resource bindings must be authored to follow a particular order
// depending on the shader format.
//
// For SPIR-V shaders, use the following resource sets:
//
// - 0: Sampled textures, followed by read-only storage textures, followed by
// read-only storage buffers
// - 1: Read-write storage textures, followed by read-write storage buffers
// - 2: Uniform buffers
//
// For DXBC and DXIL shaders, use the following register order:
//
// - (t[n], space0): Sampled textures, followed by read-only storage textures,
// followed by read-only storage buffers
// - (u[n], space1): Read-write storage textures, followed by read-write
// storage buffers
// - (b[n], space2): Uniform buffers
//
// For MSL/metallib, use the following order:
//
// - [[buffer]]: Uniform buffers, followed by read-only storage buffers,
// followed by read-write storage buffers
// - [[texture]]: Sampled textures, followed by read-only storage textures,
// followed by read-write storage textures
//
// There are optional properties that can be provided through `props`. These
// are the supported properties:
//
// - `SDL_PROP_GPU_COMPUTEPIPELINE_CREATE_NAME_STRING`: a name that can be
// displayed in debugging tools.
//
// device: a GPU Context.
//
// createinfo: a struct describing the state of the compute pipeline to
// create.
//
// Returns a compute pipeline object on success, or NULL on failure; call
// SDL_GetError() for more information.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_CreateGPUComputePipeline
func (device *GPUDevice) CreateComputePipeline(createinfo *GPUComputePipelineCreateInfo) (*GPUComputePipeline, error) {
	info := &C.SDL_GPUComputePipelineCreateInfo{
		code_size:                      C.size_t(len(createinfo.Code)),
		code:                           (*C.Uint8)(unsafe.SliceData(createinfo.Code)),
		entrypoint:                     tmpstring(createinfo.Entrypoint),
		format:                         C.SDL_GPUShaderFormat(createinfo.Format),
		num_samplers:                   C.Uint32(createinfo.NumSamplers),
		num_readonly_storage_textures:  C.Uint32(createinfo.NumReadonlyStorageTextures),
		num_readonly_storage_buffers:   C.Uint32(createinfo.NumReadonlyStorageBuffers),
		num_readwrite_storage_textures: C.Uint32(createinfo.NumReadwriteStorageTextures),
		num_readwrite_storage_buffers:  C.Uint32(createinfo.NumReadwriteStorageBuffers),
		num_uniform_buffers:            C.Uint32(createinfo.NumUniformBuffers),
		threadcount_x:                  C.Uint32(createinfo.ThreadcountX),
		threadcount_y:                  C.Uint32(createinfo.ThreadcountY),
		threadcount_z:                  C.Uint32(createinfo.ThreadcountZ),
		props:                          C.SDL_PropertiesID(createinfo.Props),
	}
	var pinner runtime.Pinner
	pinner.Pin(info.code)
	pipeline := (*GPUComputePipeline)(C.SDL_CreateGPUComputePipeline((*C.SDL_GPUDevice)(device), info))
	pinner.Unpin()
	if pipeline == nil {
		return nil, getError()
	}
	return pipeline, nil
}

const PropGPUComputepipelineCreateNameString = "SDL.gpu.computepipeline.create.name"

// Creates a pipeline object to be used in a graphics workflow.
//
// There are optional properties that can be provided through `props`. These
// are the supported properties:
//
// - `SDL_PROP_GPU_GRAPHICSPIPELINE_CREATE_NAME_STRING`: a name that can be
// displayed in debugging tools.
//
// device: a GPU Context.
//
// createinfo: a struct describing the state of the graphics pipeline to
// create.
//
// Returns a graphics pipeline object on success, or NULL on failure; call
// SDL_GetError() for more information.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_CreateGPUGraphicsPipeline
func (device *GPUDevice) CreateGraphicsPipeline(createinfo *GPUGraphicsPipelineCreateInfo) (*GPUGraphicsPipeline, error) {
	vertexBufferDescriptions := make([]C.SDL_GPUVertexBufferDescription, len(createinfo.VertexInputState.VertexBufferDescriptions))
	for i, d := range createinfo.VertexInputState.VertexBufferDescriptions {
		vertexBufferDescriptions[i] = C.SDL_GPUVertexBufferDescription{
			slot:               C.Uint32(d.Slot),
			pitch:              C.Uint32(d.Pitch),
			input_rate:         C.SDL_GPUVertexInputRate(d.InputRate),
			instance_step_rate: C.Uint32(d.InstanceStepRate),
		}
	}
	vertexAttributes := make([]C.SDL_GPUVertexAttribute, len(createinfo.VertexInputState.VertexAttributes))
	for i, a := range createinfo.VertexInputState.VertexAttributes {
		vertexAttributes[i] = C.SDL_GPUVertexAttribute{
			location:    C.Uint32(a.Location),
			buffer_slot: C.Uint32(a.BufferSlot),
			format:      C.SDL_GPUVertexElementFormat(a.Format),
			offset:      C.Uint32(a.Offset),
		}
	}
	colorTargetDescriptions := make([]C.SDL_GPUColorTargetDescription, len(createinfo.TargetInfo.ColorTargetDescriptions))
	for i, d := range createinfo.TargetInfo.ColorTargetDescriptions {
		colorTargetDescriptions[i] = C.SDL_GPUColorTargetDescription{
			format: C.SDL_GPUTextureFormat(d.Format),
			blend_state: C.SDL_GPUColorTargetBlendState{
				src_color_blendfactor:   C.SDL_GPUBlendFactor(d.BlendState.SrcColorBlendfactor),
				dst_color_blendfactor:   C.SDL_GPUBlendFactor(d.BlendState.DstColorBlendfactor),
				color_blend_op:          C.SDL_GPUBlendOp(d.BlendState.ColorBlendOp),
				src_alpha_blendfactor:   C.SDL_GPUBlendFactor(d.BlendState.SrcAlphaBlendfactor),
				dst_alpha_blendfactor:   C.SDL_GPUBlendFactor(d.BlendState.DstAlphaBlendfactor),
				alpha_blend_op:          C.SDL_GPUBlendOp(d.BlendState.AlphaBlendOp),
				color_write_mask:        C.SDL_GPUColorComponentFlags(d.BlendState.ColorWriteMask),
				enable_blend:            C.bool(d.BlendState.EnableBlend),
				enable_color_write_mask: C.bool(d.BlendState.EnableColorWriteMask),
			},
		}
	}
	info := &C.SDL_GPUGraphicsPipelineCreateInfo{
		vertex_shader:   (*C.SDL_GPUShader)(createinfo.VertexShader),
		fragment_shader: (*C.SDL_GPUShader)(createinfo.FragmentShader),
		vertex_input_state: C.SDL_GPUVertexInputState{
			vertex_buffer_descriptions: unsafe.SliceData(vertexBufferDescriptions),
			num_vertex_buffers:         C.Uint32(len(vertexBufferDescriptions)),
			vertex_attributes:          unsafe.SliceData(vertexAttributes),
			num_vertex_attributes:      C.Uint32(len(vertexAttributes)),
		},
		primitive_type: C.SDL_GPUPrimitiveType(createinfo.PrimitiveType),
		rasterizer_state: C.SDL_GPURasterizerState{
			fill_mode:                  C.SDL_GPUFillMode(createinfo.RasterizerState.FillMode),
			cull_mode:                  C.SDL_GPUCullMode(createinfo.RasterizerState.CullMode),
			front_face:                 C.SDL_GPUFrontFace(createinfo.RasterizerState.FrontFace),
			depth_bias_constant_factor: C.float(createinfo.RasterizerState.DepthBiasConstantFactor),
			depth_bias_clamp:           C.float(createinfo.RasterizerState.DepthBiasClamp),
			depth_bias_slope_factor:    C.float(createinfo.RasterizerState.DepthBiasSlopeFactor),
			enable_depth_bias:          C.bool(createinfo.RasterizerState.EnableDepthBias),
			enable_depth_clip:          C.bool(createinfo.RasterizerState.EnableDepthClip),
		},
		multisample_state: C.SDL_GPUMultisampleState{
			sample_count: C.SDL_GPUSampleCount(createinfo.MultisampleState.SampleCount),
			sample_mask:  C.Uint32(createinfo.MultisampleState.SampleMask),
			enable_mask:  C.bool(createinfo.MultisampleState.EnableMask),
		},
		depth_stencil_state: C.SDL_GPUDepthStencilState{
			compare_op: C.SDL_GPUCompareOp(createinfo.DepthStencilState.CompareOp),
			back_stencil_state: C.SDL_GPUStencilOpState{
				fail_op:       C.SDL_GPUStencilOp(createinfo.DepthStencilState.BackStencilState.FailOp),
				pass_op:       C.SDL_GPUStencilOp(createinfo.DepthStencilState.BackStencilState.PassOp),
				depth_fail_op: C.SDL_GPUStencilOp(createinfo.DepthStencilState.BackStencilState.DepthFailOp),
				compare_op:    C.SDL_GPUCompareOp(createinfo.DepthStencilState.BackStencilState.CompareOp),
			},
			front_stencil_state: C.SDL_GPUStencilOpState{
				fail_op:       C.SDL_GPUStencilOp(createinfo.DepthStencilState.FrontStencilState.FailOp),
				pass_op:       C.SDL_GPUStencilOp(createinfo.DepthStencilState.FrontStencilState.PassOp),
				depth_fail_op: C.SDL_GPUStencilOp(createinfo.DepthStencilState.FrontStencilState.DepthFailOp),
				compare_op:    C.SDL_GPUCompareOp(createinfo.DepthStencilState.FrontStencilState.CompareOp),
			},
			compare_mask:        C.Uint8(createinfo.DepthStencilState.CompareMask),
			write_mask:          C.Uint8(createinfo.DepthStencilState.WriteMask),
			enable_depth_test:   C.bool(createinfo.DepthStencilState.EnableDepthTest),
			enable_depth_write:  C.bool(createinfo.DepthStencilState.EnableDepthWrite),
			enable_stencil_test: C.bool(createinfo.DepthStencilState.EnableStencilTest),
		},
		target_info: C.SDL_GPUGraphicsPipelineTargetInfo{
			color_target_descriptions: unsafe.SliceData(colorTargetDescriptions),
			num_color_targets:         C.Uint32(len(colorTargetDescriptions)),
			depth_stencil_format:      C.SDL_GPUTextureFormat(createinfo.TargetInfo.DepthStencilFormat),
			has_depth_stencil_target:  C.bool(createinfo.TargetInfo.HasDepthStencilTarget),
		},
		props: C.SDL_PropertiesID(createinfo.Props),
	}
	var pinner runtime.Pinner
	pinner.Pin(info.vertex_input_state.vertex_buffer_descriptions)
	pinner.Pin(info.vertex_input_state.vertex_attributes)
	pinner.Pin(info.target_info.color_target_descriptions)
	pipeline := (*GPUGraphicsPipeline)(C.SDL_CreateGPUGraphicsPipeline((*C.SDL_GPUDevice)(device), info))
	pinner.Unpin()
	if pipeline == nil {
		return nil, getError()
	}
	return pipeline, nil
}

const PropGPUGraphicspipelineCreateNameString = "SDL.gpu.graphicspipeline.create.name"

// Creates a sampler object to be used when binding textures in a graphics
// workflow.
//
// There are optional properties that can be provided through `props`. These
// are the supported properties:
//
// - `SDL_PROP_GPU_SAMPLER_CREATE_NAME_STRING`: a name that can be displayed
// in debugging tools.
//
// device: a GPU Context.
//
// createinfo: a struct describing the state of the sampler to create.
//
// Returns a sampler object on success, or NULL on failure; call
// SDL_GetError() for more information.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_CreateGPUSampler
func (device *GPUDevice) CreateSampler(createinfo *GPUSamplerCreateInfo) (*GPUSampler, error) {
	sampler := (*GPUSampler)(C.SDL_CreateGPUSampler((*C.SDL_GPUDevice)(device), &C.SDL_GPUSamplerCreateInfo{
		min_filter:        C.SDL_GPUFilter(createinfo.MinFilter),
		mag_filter:        C.SDL_GPUFilter(createinfo.MagFilter),
		mipmap_mode:       C.SDL_GPUSamplerMipmapMode(createinfo.MipmapMode),
		address_mode_u:    C.SDL_GPUSamplerAddressMode(createinfo.AddressModeU),
		address_mode_v:    C.SDL_GPUSamplerAddressMode(createinfo.AddressModeV),
		address_mode_w:    C.SDL_GPUSamplerAddressMode(createinfo.AddressModeW),
		mip_lod_bias:      C.float(createinfo.MipLodBias),
		max_anisotropy:    C.float(createinfo.MaxAnisotropy),
		compare_op:        C.SDL_GPUCompareOp(createinfo.CompareOp),
		min_lod:           C.float(createinfo.MinLod),
		max_lod:           C.float(createinfo.MaxLod),
		enable_anisotropy: C.bool(createinfo.EnableAnisotropy),
		enable_compare:    C.bool(createinfo.EnableCompare),
		props:             C.SDL_PropertiesID(createinfo.Props),
	}))
	if sampler == nil {
		return nil, getError()
	}
	return sampler, nil
}

const PropGPUSamplerCreateNameString = "SDL.gpu.sampler.create.name"

// Creates a shader to be used when creating a graphics pipeline.
//
// Shader resource bindings must be authored to follow a particular order
// depending on the shader format.
//
// For SPIR-V shaders, use the following resource sets:
//
// For vertex shaders:
//
// - 0: Sampled textures, followed by storage textures, followed by storage
// buffers
// - 1: Uniform buffers
//
// For fragment shaders:
//
// - 2: Sampled textures, followed by storage textures, followed by storage
// buffers
// - 3: Uniform buffers
//
// For DXBC and DXIL shaders, use the following register order:
//
// For vertex shaders:
//
// - (t[n], space0): Sampled textures, followed by storage textures, followed
// by storage buffers
// - (s[n], space0): Samplers with indices corresponding to the sampled
// textures
// - (b[n], space1): Uniform buffers
//
// For pixel shaders:
//
// - (t[n], space2): Sampled textures, followed by storage textures, followed
// by storage buffers
// - (s[n], space2): Samplers with indices corresponding to the sampled
// textures
// - (b[n], space3): Uniform buffers
//
// For MSL/metallib, use the following order:
//
// - [[texture]]: Sampled textures, followed by storage textures
// - [[sampler]]: Samplers with indices corresponding to the sampled textures
// - [[buffer]]: Uniform buffers, followed by storage buffers. Vertex buffer 0
// is bound at [[buffer(14)]], vertex buffer 1 at [[buffer(15)]], and so on.
// Rather than manually authoring vertex buffer indices, use the
// [[stage_in]] attribute which will automatically use the vertex input
// information from the SDL_GPUGraphicsPipeline.
//
// Shader semantics other than system-value semantics do not matter in D3D12
// and for ease of use the SDL implementation assumes that non system-value
// semantics will all be TEXCOORD. If you are using HLSL as the shader source
// language, your vertex semantics should start at TEXCOORD0 and increment
// like so: TEXCOORD1, TEXCOORD2, etc. If you wish to change the semantic
// prefix to something other than TEXCOORD you can use
// SDL_PROP_GPU_DEVICE_CREATE_D3D12_SEMANTIC_NAME_STRING with
// SDL_CreateGPUDeviceWithProperties().
//
// There are optional properties that can be provided through `props`. These
// are the supported properties:
//
// - `SDL_PROP_GPU_SHADER_CREATE_NAME_STRING`: a name that can be displayed in
// debugging tools.
//
// device: a GPU Context.
//
// createinfo: a struct describing the state of the shader to create.
//
// Returns a shader object on success, or NULL on failure; call
// SDL_GetError() for more information.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_CreateGPUShader
func (device *GPUDevice) CreateShader(createinfo *GPUShaderCreateInfo) (*GPUShader, error) {
	info := &C.SDL_GPUShaderCreateInfo{
		code_size:            C.size_t(len(createinfo.Code)),
		code:                 (*C.Uint8)(unsafe.SliceData(createinfo.Code)),
		entrypoint:           tmpstring(createinfo.Entrypoint),
		format:               C.SDL_GPUShaderFormat(createinfo.Format),
		stage:                C.SDL_GPUShaderStage(createinfo.Stage),
		num_samplers:         C.Uint32(createinfo.NumSamplers),
		num_storage_textures: C.Uint32(createinfo.NumStorageTextures),
		num_storage_buffers:  C.Uint32(createinfo.NumStorageBuffers),
		num_uniform_buffers:  C.Uint32(createinfo.NumUniformBuffers),
		props:                C.SDL_PropertiesID(createinfo.Props),
	}
	var pinner runtime.Pinner
	pinner.Pin(info.code)
	shader := (*GPUShader)(C.SDL_CreateGPUShader((*C.SDL_GPUDevice)(device), info))
	pinner.Unpin()
	if shader == nil {
		return nil, getError()
	}
	return shader, nil
}

const PropGPUShaderCreateNameString = "SDL.gpu.shader.create.name"

// Creates a texture object to be used in graphics or compute workflows.
//
// The contents of this texture are undefined until data is written to the
// texture.
//
// Note that certain combinations of usage flags are invalid. For example, a
// texture cannot have both the SAMPLER and GRAPHICS_STORAGE_READ flags.
//
// If you request a sample count higher than the hardware supports, the
// implementation will automatically fall back to the highest available sample
// count.
//
// There are optional properties that can be provided through
// SDL_GPUTextureCreateInfo's `props`. These are the supported properties:
//
// - `SDL_PROP_GPU_TEXTURE_CREATE_D3D12_CLEAR_R_FLOAT`: (Direct3D 12 only) if
// the texture usage is SDL_GPU_TEXTUREUSAGE_COLOR_TARGET, clear the texture
// to a color with this red intensity. Defaults to zero.
// - `SDL_PROP_GPU_TEXTURE_CREATE_D3D12_CLEAR_G_FLOAT`: (Direct3D 12 only) if
// the texture usage is SDL_GPU_TEXTUREUSAGE_COLOR_TARGET, clear the texture
// to a color with this green intensity. Defaults to zero.
// - `SDL_PROP_GPU_TEXTURE_CREATE_D3D12_CLEAR_B_FLOAT`: (Direct3D 12 only) if
// the texture usage is SDL_GPU_TEXTUREUSAGE_COLOR_TARGET, clear the texture
// to a color with this blue intensity. Defaults to zero.
// - `SDL_PROP_GPU_TEXTURE_CREATE_D3D12_CLEAR_A_FLOAT`: (Direct3D 12 only) if
// the texture usage is SDL_GPU_TEXTUREUSAGE_COLOR_TARGET, clear the texture
// to a color with this alpha intensity. Defaults to zero.
// - `SDL_PROP_GPU_TEXTURE_CREATE_D3D12_CLEAR_DEPTH_FLOAT`: (Direct3D 12 only)
// if the texture usage is SDL_GPU_TEXTUREUSAGE_DEPTH_STENCIL_TARGET, clear
// the texture to a depth of this value. Defaults to zero.
// - `SDL_PROP_GPU_TEXTURE_CREATE_D3D12_CLEAR_STENCIL_UINT8`: (Direct3D 12
// only) if the texture usage is SDL_GPU_TEXTUREUSAGE_DEPTH_STENCIL_TARGET,
// clear the texture to a stencil of this value. Defaults to zero.
// - `SDL_PROP_GPU_TEXTURE_CREATE_NAME_STRING`: a name that can be displayed
// in debugging tools.
//
// device: a GPU Context.
//
// createinfo: a struct describing the state of the texture to create.
//
// Returns a texture object on success, or NULL on failure; call
// SDL_GetError() for more information.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_CreateGPUTexture
func (device *GPUDevice) CreateTexture(createinfo *GPUTextureCreateInfo) (*GPUTexture, error) {
	texture := (*GPUTexture)(C.SDL_CreateGPUTexture((*C.SDL_GPUDevice)(device), &C.SDL_GPUTextureCreateInfo{
		_type:                C.SDL_GPUTextureType(createinfo.Type),
		format:               C.SDL_GPUTextureFormat(createinfo.Format),
		usage:                C.SDL_GPUTextureUsageFlags(createinfo.Usage),
		width:                C.Uint32(createinfo.Width),
		height:               C.Uint32(createinfo.Height),
		layer_count_or_depth: C.Uint32(createinfo.LayerCountOrDepth),
		num_levels:           C.Uint32(createinfo.NumLevels),
		sample_count:         C.SDL_GPUSampleCount(createinfo.SampleCount),
		props:                C.SDL_PropertiesID(createinfo.Props),
	}))
	if texture == nil {
		return nil, getError()
	}
	return texture, nil
}

const PropGPUTextureCreateD3D12ClearRFloat = "SDL.gpu.texture.create.d3d12.clear.r"
const PropGPUTextureCreateD3D12ClearGFloat = "SDL.gpu.texture.create.d3d12.clear.g"
const PropGPUTextureCreateD3D12ClearBFloat = "SDL.gpu.texture.create.d3d12.clear.b"
const PropGPUTextureCreateD3D12ClearAFloat = "SDL.gpu.texture.create.d3d12.clear.a"
const PropGPUTextureCreateD3D12ClearDepthFloat = "SDL.gpu.texture.create.d3d12.clear.depth"
const PropGPUTextureCreateD3D12ClearStencilUINT8 = "SDL.gpu.texture.create.d3d12.clear.stencil"
const PropGPUTextureCreateNameString = "SDL.gpu.texture.create.name"

// Creates a buffer object to be used in graphics or compute workflows.
//
// The contents of this buffer are undefined until data is written to the
// buffer.
//
// Note that certain combinations of usage flags are invalid. For example, a
// buffer cannot have both the VERTEX and INDEX flags.
//
// For better understanding of underlying concepts and memory management with
// SDL GPU API, you may refer
// [this blog post](https://moonside.games/posts/sdl-gpu-concepts-cycling/)
// .
//
// There are optional properties that can be provided through `props`. These
// are the supported properties:
//
// - `SDL_PROP_GPU_BUFFER_CREATE_NAME_STRING`: a name that can be displayed in
// debugging tools.
//
// device: a GPU Context.
//
// createinfo: a struct describing the state of the buffer to create.
//
// Returns a buffer object on success, or NULL on failure; call
// SDL_GetError() for more information.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_CreateGPUBuffer
func (device *GPUDevice) CreateBuffer(createinfo *GPUBufferCreateInfo) (*GPUBuffer, error) {
	buffer := (*GPUBuffer)(C.SDL_CreateGPUBuffer((*C.SDL_GPUDevice)(device), &C.SDL_GPUBufferCreateInfo{
		usage: C.SDL_GPUBufferUsageFlags(createinfo.Usage),
		size:  C.Uint32(createinfo.Size),
		props: C.SDL_PropertiesID(createinfo.Props),
	}))
	if buffer == nil {
		return nil, getError()
	}
	return buffer, nil
}

const PropGPUBufferCreateNameString = "SDL.gpu.buffer.create.name"

// Creates a transfer buffer to be used when uploading to or downloading from
// graphics resources.
//
// Download buffers can be particularly expensive to create, so it is good
// practice to reuse them if data will be downloaded regularly.
//
// There are optional properties that can be provided through `props`. These
// are the supported properties:
//
// - `SDL_PROP_GPU_TRANSFERBUFFER_CREATE_NAME_STRING`: a name that can be
// displayed in debugging tools.
//
// device: a GPU Context.
//
// createinfo: a struct describing the state of the transfer buffer to
// create.
//
// Returns a transfer buffer on success, or NULL on failure; call
// SDL_GetError() for more information.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_CreateGPUTransferBuffer
func (device *GPUDevice) CreateTransferBuffer(createinfo *GPUTransferBufferCreateInfo) (*GPUTransferBuffer, error) {
	buffer := C.SDL_CreateGPUTransferBuffer((*C.SDL_GPUDevice)(device), &C.SDL_GPUTransferBufferCreateInfo{
		usage: C.SDL_GPUTransferBufferUsage(createinfo.Usage),
		size:  C.Uint32(createinfo.Size),
		props: C.SDL_PropertiesID(createinfo.Props),
	})
	if buffer == nil {
		return nil, getError()
	}
	return &GPUTransferBuffer{buffer, createinfo.Size, nil}, nil
}

const PropGPUTransferbufferCreateNameString = "SDL.gpu.transferbuffer.create.name"

// Sets an arbitrary string constant to label a buffer.
//
// You should use SDL_PROP_GPU_BUFFER_CREATE_NAME_STRING with
// SDL_CreateGPUBuffer instead of this function to avoid thread safety issues.
//
// device: a GPU Context.
//
// buffer: a buffer to attach the name to.
//
// text: a UTF-8 string constant to mark as the name of the buffer.
//
// This function is not thread safe, you must make sure the
// buffer is not simultaneously used by any other thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetGPUBufferName
func (device *GPUDevice) SetBufferName(buffer *GPUBuffer, text string) {
	C.SDL_SetGPUBufferName((*C.SDL_GPUDevice)(device), (*C.SDL_GPUBuffer)(buffer), tmpstring(text))
}

// Sets an arbitrary string constant to label a texture.
//
// You should use SDL_PROP_GPU_TEXTURE_CREATE_NAME_STRING with
// SDL_CreateGPUTexture instead of this function to avoid thread safety
// issues.
//
// device: a GPU Context.
//
// texture: a texture to attach the name to.
//
// text: a UTF-8 string constant to mark as the name of the texture.
//
// This function is not thread safe, you must make sure the
// texture is not simultaneously used by any other thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetGPUTextureName
func (device *GPUDevice) SetTextureName(texture *GPUTexture, text string) {
	C.SDL_SetGPUTextureName((*C.SDL_GPUDevice)(device), (*C.SDL_GPUTexture)(texture), tmpstring(text))
}

// Inserts an arbitrary string label into the command buffer callstream.
//
// Useful for debugging.
//
// command_buffer: a command buffer.
//
// text: a UTF-8 string constant to insert as the label.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_InsertGPUDebugLabel
func (cmd *GPUCommandBuffer) InsertDebugLabel(text string) {
	C.SDL_InsertGPUDebugLabel((*C.SDL_GPUCommandBuffer)(cmd), tmpstring(text))
}

// Begins a debug group with an arbitary name.
//
// Used for denoting groups of calls when viewing the command buffer
// callstream in a graphics debugging tool.
//
// Each call to SDL_PushGPUDebugGroup must have a corresponding call to
// SDL_PopGPUDebugGroup.
//
// On some backends (e.g. Metal), pushing a debug group during a
// render/blit/compute pass will create a group that is scoped to the native
// pass rather than the command buffer. For best results, if you push a debug
// group during a pass, always pop it in the same pass.
//
// command_buffer: a command buffer.
//
// name: a UTF-8 string constant that names the group.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_PushGPUDebugGroup
func (cmd *GPUCommandBuffer) PushDebugGroup(name string) {
	C.SDL_PushGPUDebugGroup((*C.SDL_GPUCommandBuffer)(cmd), tmpstring(name))
}

// Ends the most-recently pushed debug group.
//
// command_buffer: a command buffer.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_PopGPUDebugGroup
func (cmd *GPUCommandBuffer) PopDebugGroup() {
	C.SDL_PopGPUDebugGroup((*C.SDL_GPUCommandBuffer)(cmd))
}

// Frees the given texture as soon as it is safe to do so.
//
// You must not reference the texture after calling this function.
//
// device: a GPU context.
//
// texture: a texture to be destroyed.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_ReleaseGPUTexture
func (device *GPUDevice) ReleaseTexture(texture *GPUTexture) {
	C.SDL_ReleaseGPUTexture((*C.SDL_GPUDevice)(device), (*C.SDL_GPUTexture)(texture))
}

// Frees the given sampler as soon as it is safe to do so.
//
// You must not reference the sampler after calling this function.
//
// device: a GPU context.
//
// sampler: a sampler to be destroyed.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_ReleaseGPUSampler
func (device *GPUDevice) ReleaseSampler(sampler *GPUSampler) {
	C.SDL_ReleaseGPUSampler((*C.SDL_GPUDevice)(device), (*C.SDL_GPUSampler)(sampler))
}

// Frees the given buffer as soon as it is safe to do so.
//
// You must not reference the buffer after calling this function.
//
// device: a GPU context.
//
// buffer: a buffer to be destroyed.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_ReleaseGPUBuffer
func (device *GPUDevice) ReleaseBuffer(buffer *GPUBuffer) {
	C.SDL_ReleaseGPUBuffer((*C.SDL_GPUDevice)(device), (*C.SDL_GPUBuffer)(buffer))
}

// Frees the given transfer buffer as soon as it is safe to do so.
//
// You must not reference the transfer buffer after calling this function.
//
// device: a GPU context.
//
// transfer_buffer: a transfer buffer to be destroyed.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_ReleaseGPUTransferBuffer
func (device *GPUDevice) ReleaseTransferBuffer(transferBuffer *GPUTransferBuffer) {
	C.SDL_ReleaseGPUTransferBuffer((*C.SDL_GPUDevice)(device), transferBuffer.internal)
	if transferBuffer.mapping != nil {
		transferBuffer.mapping.p = nil
	}
}

// Frees the given compute pipeline as soon as it is safe to do so.
//
// You must not reference the compute pipeline after calling this function.
//
// device: a GPU context.
//
// compute_pipeline: a compute pipeline to be destroyed.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_ReleaseGPUComputePipeline
func (device *GPUDevice) ReleaseComputePipeline(computePipeline *GPUComputePipeline) {
	C.SDL_ReleaseGPUComputePipeline((*C.SDL_GPUDevice)(device), (*C.SDL_GPUComputePipeline)(computePipeline))
}

// Frees the given shader as soon as it is safe to do so.
//
// You must not reference the shader after calling this function.
//
// device: a GPU context.
//
// shader: a shader to be destroyed.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_ReleaseGPUShader
func (device *GPUDevice) ReleaseShader(shader *GPUShader) {
	C.SDL_ReleaseGPUShader((*C.SDL_GPUDevice)(device), (*C.SDL_GPUShader)(shader))
}

// Frees the given graphics pipeline as soon as it is safe to do so.
//
// You must not reference the graphics pipeline after calling this function.
//
// device: a GPU context.
//
// graphics_pipeline: a graphics pipeline to be destroyed.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_ReleaseGPUGraphicsPipeline
func (device *GPUDevice) ReleaseGraphicsPipeline(graphicsPipeline *GPUGraphicsPipeline) {
	C.SDL_ReleaseGPUGraphicsPipeline((*C.SDL_GPUDevice)(device), (*C.SDL_GPUGraphicsPipeline)(graphicsPipeline))
}

// Acquire a command buffer.
//
// This command buffer is managed by the implementation and should not be
// freed by the user. The command buffer may only be used on the thread it was
// acquired on. The command buffer should be submitted on the thread it was
// acquired on.
//
// It is valid to acquire multiple command buffers on the same thread at once.
// In fact a common design pattern is to acquire two command buffers per frame
// where one is dedicated to render and compute passes and the other is
// dedicated to copy passes and other preparatory work such as generating
// mipmaps. Interleaving commands between the two command buffers reduces the
// total amount of passes overall which improves rendering performance.
//
// device: a GPU context.
//
// Returns a command buffer, or NULL on failure; call SDL_GetError() for more
// information.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_AcquireGPUCommandBuffer
func (device *GPUDevice) AcquireCommandBuffer() (*GPUCommandBuffer, error) {
	cmd := (*GPUCommandBuffer)(C.SDL_AcquireGPUCommandBuffer((*C.SDL_GPUDevice)(device)))
	if cmd == nil {
		return nil, getError()
	}
	return cmd, nil
}

// Pushes data to a vertex uniform slot on the command buffer.
//
// Subsequent draw calls will use this uniform data.
//
// command_buffer: a command buffer.
//
// slot_index: the vertex uniform slot to push data to.
//
// data: client data to write.
//
// length: the length of the data to write.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_PushGPUVertexUniformData
func (cmd *GPUCommandBuffer) PushVertexUniformData(slotIndex int, data []byte) {
	C.SDL_PushGPUVertexUniformData((*C.SDL_GPUCommandBuffer)(cmd), (C.Uint32)(slotIndex), unsafe.Pointer(unsafe.SliceData(data)), (C.Uint32)(len(data)))
}

// Pushes data to a fragment uniform slot on the command buffer.
//
// Subsequent draw calls will use this uniform data.
//
// command_buffer: a command buffer.
//
// slot_index: the fragment uniform slot to push data to.
//
// data: client data to write.
//
// length: the length of the data to write.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_PushGPUFragmentUniformData
func (cmd *GPUCommandBuffer) PushFragmentUniformData(slotIndex int, data []byte) {
	C.SDL_PushGPUFragmentUniformData((*C.SDL_GPUCommandBuffer)(cmd), (C.Uint32)(slotIndex), unsafe.Pointer(unsafe.SliceData(data)), (C.Uint32)(len(data)))
}

// Pushes data to a uniform slot on the command buffer.
//
// Subsequent draw calls will use this uniform data.
//
// command_buffer: a command buffer.
//
// slot_index: the uniform slot to push data to.
//
// data: client data to write.
//
// length: the length of the data to write.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_PushGPUComputeUniformData
func (cmd *GPUCommandBuffer) PushComputeUniformData(slotIndex int, data []byte) {
	C.SDL_PushGPUComputeUniformData((*C.SDL_GPUCommandBuffer)(cmd), (C.Uint32)(slotIndex), unsafe.Pointer(unsafe.SliceData(data)), (C.Uint32)(len(data)))
}

// Begins a render pass on a command buffer.
//
// A render pass consists of a set of texture subresources (or depth slices in
// the 3D texture case) which will be rendered to during the render pass,
// along with corresponding clear values and load/store operations. All
// operations related to graphics pipelines must take place inside of a render
// pass. A default viewport and scissor state are automatically set when this
// is called. You cannot begin another render pass, or begin a compute pass or
// copy pass until you have ended the render pass.
//
// command_buffer: a command buffer.
//
// color_target_infos: an array of texture subresources with
// corresponding clear values and load/store ops.
//
// num_color_targets: the number of color targets in the
// color_target_infos array.
//
// depth_stencil_target_info: a texture subresource with corresponding
// clear value and load/store ops, may be
// NULL.
//
// Returns a render pass handle.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_BeginGPURenderPass
func (cmd *GPUCommandBuffer) BeginRenderPass(colorTargetInfos []GPUColorTargetInfo, depthStencilTargetInfo *GPUDepthStencilTargetInfo) *GPURenderPass {
	colorTargets := make([]C.SDL_GPUColorTargetInfo, len(colorTargetInfos))
	for i, t := range colorTargetInfos {
		colorTargets[i] = C.SDL_GPUColorTargetInfo{
			texture:               (*C.SDL_GPUTexture)(t.Texture),
			mip_level:             C.Uint32(t.MipLevel),
			layer_or_depth_plane:  C.Uint32(t.LayerOrDepthPlane),
			clear_color:           C.SDL_FColor{C.float(t.ClearColor.R), C.float(t.ClearColor.G), C.float(t.ClearColor.B), C.float(t.ClearColor.A)},
			load_op:               C.SDL_GPULoadOp(t.LoadOp),
			store_op:              C.SDL_GPUStoreOp(t.StoreOp),
			resolve_texture:       (*C.SDL_GPUTexture)(t.ResolveTexture),
			resolve_mip_level:     C.Uint32(t.ResolveMipLevel),
			resolve_layer:         C.Uint32(t.ResolveLayer),
			cycle:                 C.bool(t.Cycle),
			cycle_resolve_texture: C.bool(t.CycleResolveTexture),
		}
	}
	var depthStencilTarget *C.SDL_GPUDepthStencilTargetInfo
	if depthStencilTargetInfo != nil {
		depthStencilTarget = &C.SDL_GPUDepthStencilTargetInfo{
			texture:          (*C.SDL_GPUTexture)(depthStencilTargetInfo.Texture),
			clear_depth:      C.float(depthStencilTargetInfo.ClearDepth),
			load_op:          C.SDL_GPULoadOp(depthStencilTargetInfo.LoadOp),
			store_op:         C.SDL_GPUStoreOp(depthStencilTargetInfo.StoreOp),
			stencil_load_op:  C.SDL_GPULoadOp(depthStencilTargetInfo.StencilLoadOp),
			stencil_store_op: C.SDL_GPUStoreOp(depthStencilTargetInfo.StencilStoreOp),
			cycle:            C.bool(depthStencilTargetInfo.Cycle),
			clear_stencil:    C.Uint8(depthStencilTargetInfo.ClearStencil),
		}
	}
	return (*GPURenderPass)(C.SDL_BeginGPURenderPass((*C.SDL_GPUCommandBuffer)(cmd), unsafe.SliceData(colorTargets), (C.Uint32)(len(colorTargets)), depthStencilTarget))
}

// Binds a graphics pipeline on a render pass to be used in rendering.
//
// A graphics pipeline must be bound before making any draw calls.
//
// render_pass: a render pass handle.
//
// graphics_pipeline: the graphics pipeline to bind.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_BindGPUGraphicsPipeline
func (pass *GPURenderPass) BindPipeline(graphicsPipeline *GPUGraphicsPipeline) {
	C.SDL_BindGPUGraphicsPipeline((*C.SDL_GPURenderPass)(pass), (*C.SDL_GPUGraphicsPipeline)(graphicsPipeline))
}

// Sets the current viewport state on a command buffer.
//
// render_pass: a render pass handle.
//
// viewport: the viewport to set.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetGPUViewport
func (pass *GPURenderPass) SetViewport(viewport *GPUViewport) {
	C.SDL_SetGPUViewport((*C.SDL_GPURenderPass)(pass), &C.SDL_GPUViewport{C.float(viewport.X), C.float(viewport.Y), C.float(viewport.W), C.float(viewport.H), C.float(viewport.MinDepth), C.float(viewport.MaxDepth)})
}

// Sets the current scissor state on a command buffer.
//
// render_pass: a render pass handle.
//
// scissor: the scissor area to set.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetGPUScissor
func (pass *GPURenderPass) SetScissor(scissor Rect) {
	C.SDL_SetGPUScissor((*C.SDL_GPURenderPass)(pass), &C.SDL_Rect{C.int(scissor.X), C.int(scissor.Y), C.int(scissor.W), C.int(scissor.H)})
}

// Sets the current blend constants on a command buffer.
//
// render_pass: a render pass handle.
//
// blend_constants: the blend constant color.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetGPUBlendConstants
func (pass *GPURenderPass) SetBlendConstants(c FColor) {
	C.SDL_SetGPUBlendConstants((*C.SDL_GPURenderPass)(pass), C.SDL_FColor{C.float(c.R), C.float(c.G), C.float(c.B), C.float(c.A)})
}

// Sets the current stencil reference value on a command buffer.
//
// render_pass: a render pass handle.
//
// reference: the stencil reference value to set.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetGPUStencilReference
func (pass *GPURenderPass) SetStencilReference(reference byte) {
	C.SDL_SetGPUStencilReference((*C.SDL_GPURenderPass)(pass), (C.Uint8)(reference))
}

// Binds vertex buffers on a command buffer for use with subsequent draw
// calls.
//
// render_pass: a render pass handle.
//
// first_slot: the vertex buffer slot to begin binding from.
//
// bindings: an array of SDL_GPUBufferBinding structs containing vertex
// buffers and offset values.
//
// num_bindings: the number of bindings in the bindings array.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_BindGPUVertexBuffers
func (pass *GPURenderPass) BindVertexBuffers(firstSlot int, bindings []GPUBufferBinding) {
	cbindings := make([]C.SDL_GPUBufferBinding, len(bindings))
	for i, b := range bindings {
		cbindings[i] = C.SDL_GPUBufferBinding{buffer: (*C.SDL_GPUBuffer)(b.Buffer), offset: C.Uint32(b.Offset)}
	}
	C.SDL_BindGPUVertexBuffers((*C.SDL_GPURenderPass)(pass), (C.Uint32)(firstSlot), unsafe.SliceData(cbindings), (C.Uint32)(len(cbindings)))
}

// Binds an index buffer on a command buffer for use with subsequent draw
// calls.
//
// render_pass: a render pass handle.
//
// binding: a pointer to a struct containing an index buffer and offset.
//
// index_element_size: whether the index values in the buffer are 16- or
// 32-bit.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_BindGPUIndexBuffer
func (pass *GPURenderPass) BindIndexBuffer(binding *GPUBufferBinding, indexElementSize GPUIndexElementSize) {
	C.SDL_BindGPUIndexBuffer((*C.SDL_GPURenderPass)(pass), &C.SDL_GPUBufferBinding{buffer: (*C.SDL_GPUBuffer)(binding.Buffer), offset: C.Uint32(binding.Offset)}, (C.SDL_GPUIndexElementSize)(indexElementSize))
}

// Binds texture-sampler pairs for use on the vertex shader.
//
// The textures must have been created with SDL_GPU_TEXTUREUSAGE_SAMPLER.
//
// render_pass: a render pass handle.
//
// first_slot: the vertex sampler slot to begin binding from.
//
// texture_sampler_bindings: an array of texture-sampler binding
// structs.
//
// num_bindings: the number of texture-sampler pairs to bind from the
// array.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_BindGPUVertexSamplers
func (pass *GPURenderPass) BindVertexSamplers(firstSlot int, textureSamplerBindings []GPUTextureSamplerBinding) {
	C.SDL_BindGPUVertexSamplers((*C.SDL_GPURenderPass)(pass), (C.Uint32)(firstSlot), (*C.SDL_GPUTextureSamplerBinding)(unsafe.Pointer(unsafe.SliceData(textureSamplerBindings))), (C.Uint32)(len(textureSamplerBindings)))
}

// Binds storage textures for use on the vertex shader.
//
// These textures must have been created with
// SDL_GPU_TEXTUREUSAGE_GRAPHICS_STORAGE_READ.
//
// render_pass: a render pass handle.
//
// first_slot: the vertex storage texture slot to begin binding from.
//
// storage_textures: an array of storage textures.
//
// num_bindings: the number of storage texture to bind from the array.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_BindGPUVertexStorageTextures
func (pass *GPURenderPass) BindVertexStorageTextures(firstSlot int, storageTextures []*GPUTexture) {
	C.SDL_BindGPUVertexStorageTextures((*C.SDL_GPURenderPass)(pass), (C.Uint32)(firstSlot), (**C.SDL_GPUTexture)(unsafe.Pointer(unsafe.SliceData(storageTextures))), (C.Uint32)(len(storageTextures)))
}

// Binds storage buffers for use on the vertex shader.
//
// These buffers must have been created with
// SDL_GPU_BUFFERUSAGE_GRAPHICS_STORAGE_READ.
//
// render_pass: a render pass handle.
//
// first_slot: the vertex storage buffer slot to begin binding from.
//
// storage_buffers: an array of buffers.
//
// num_bindings: the number of buffers to bind from the array.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_BindGPUVertexStorageBuffers
func (pass *GPURenderPass) BindVertexStorageBuffers(firstSlot int, storageBuffers []*GPUBuffer) {
	C.SDL_BindGPUVertexStorageBuffers((*C.SDL_GPURenderPass)(pass), (C.Uint32)(firstSlot), (**C.SDL_GPUBuffer)(unsafe.Pointer(unsafe.SliceData(storageBuffers))), (C.Uint32)(len(storageBuffers)))
}

// Binds texture-sampler pairs for use on the fragment shader.
//
// The textures must have been created with SDL_GPU_TEXTUREUSAGE_SAMPLER.
//
// render_pass: a render pass handle.
//
// first_slot: the fragment sampler slot to begin binding from.
//
// texture_sampler_bindings: an array of texture-sampler binding
// structs.
//
// num_bindings: the number of texture-sampler pairs to bind from the
// array.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_BindGPUFragmentSamplers
func (pass *GPURenderPass) BindFragmentSamplers(firstSlot int, textureSamplerBindings []GPUTextureSamplerBinding) {
	C.SDL_BindGPUFragmentSamplers((*C.SDL_GPURenderPass)(pass), (C.Uint32)(firstSlot), (*C.SDL_GPUTextureSamplerBinding)(unsafe.Pointer(unsafe.SliceData(textureSamplerBindings))), (C.Uint32)(len(textureSamplerBindings)))
}

// Binds storage textures for use on the fragment shader.
//
// These textures must have been created with
// SDL_GPU_TEXTUREUSAGE_GRAPHICS_STORAGE_READ.
//
// render_pass: a render pass handle.
//
// first_slot: the fragment storage texture slot to begin binding from.
//
// storage_textures: an array of storage textures.
//
// num_bindings: the number of storage textures to bind from the array.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_BindGPUFragmentStorageTextures
func (pass *GPURenderPass) BindFragmentStorageTextures(firstSlot int, storageTextures []*GPUTexture) {
	C.SDL_BindGPUFragmentStorageTextures((*C.SDL_GPURenderPass)(pass), (C.Uint32)(firstSlot), (**C.SDL_GPUTexture)(unsafe.Pointer(unsafe.SliceData(storageTextures))), (C.Uint32)(len(storageTextures)))
}

// Binds storage buffers for use on the fragment shader.
//
// These buffers must have been created with
// SDL_GPU_BUFFERUSAGE_GRAPHICS_STORAGE_READ.
//
// render_pass: a render pass handle.
//
// first_slot: the fragment storage buffer slot to begin binding from.
//
// storage_buffers: an array of storage buffers.
//
// num_bindings: the number of storage buffers to bind from the array.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_BindGPUFragmentStorageBuffers
func (pass *GPURenderPass) BindFragmentStorageBuffers(firstSlot int, storageBuffers []*GPUBuffer) {
	C.SDL_BindGPUFragmentStorageBuffers((*C.SDL_GPURenderPass)(pass), (C.Uint32)(firstSlot), (**C.SDL_GPUBuffer)(unsafe.Pointer(unsafe.SliceData(storageBuffers))), (C.Uint32)(len(storageBuffers)))
}

// Draws data using bound graphics state with an index buffer and instancing
// enabled.
//
// You must not call this function before binding a graphics pipeline.
//
// Note that the `first_vertex` and `first_instance` parameters are NOT
// compatible with built-in vertex/instance ID variables in shaders (for
// example, SV_VertexID); GPU APIs and shader languages do not define these
// built-in variables consistently, so if your shader depends on them, the
// only way to keep behavior consistent and portable is to always pass 0 for
// the correlating parameter in the draw calls.
//
// render_pass: a render pass handle.
//
// num_indices: the number of indices to draw per instance.
//
// num_instances: the number of instances to draw.
//
// first_index: the starting index within the index buffer.
//
// vertex_offset: value added to vertex index before indexing into the
// vertex buffer.
//
// first_instance: the ID of the first instance to draw.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_DrawGPUIndexedPrimitives
func (pass *GPURenderPass) DrawIndexedPrimitives(numIndices int, numInstances int, firstIndex int, vertexOffset int, firstInstance int) {
	C.SDL_DrawGPUIndexedPrimitives((*C.SDL_GPURenderPass)(pass), (C.Uint32)(numIndices), (C.Uint32)(numInstances), (C.Uint32)(firstIndex), (C.Sint32)(vertexOffset), (C.Uint32)(firstInstance))
}

// Draws data using bound graphics state.
//
// You must not call this function before binding a graphics pipeline.
//
// Note that the `first_vertex` and `first_instance` parameters are NOT
// compatible with built-in vertex/instance ID variables in shaders (for
// example, SV_VertexID); GPU APIs and shader languages do not define these
// built-in variables consistently, so if your shader depends on them, the
// only way to keep behavior consistent and portable is to always pass 0 for
// the correlating parameter in the draw calls.
//
// render_pass: a render pass handle.
//
// num_vertices: the number of vertices to draw.
//
// num_instances: the number of instances that will be drawn.
//
// first_vertex: the index of the first vertex to draw.
//
// first_instance: the ID of the first instance to draw.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_DrawGPUPrimitives
func (pass *GPURenderPass) DrawPrimitives(numVertices int, numInstances int, firstVertex int, firstInstance int) {
	C.SDL_DrawGPUPrimitives((*C.SDL_GPURenderPass)(pass), (C.Uint32)(numVertices), (C.Uint32)(numInstances), (C.Uint32)(firstVertex), (C.Uint32)(firstInstance))
}

// Draws data using bound graphics state and with draw parameters set from a
// buffer.
//
// The buffer must consist of tightly-packed draw parameter sets that each
// match the layout of SDL_GPUIndirectDrawCommand. You must not call this
// function before binding a graphics pipeline.
//
// render_pass: a render pass handle.
//
// buffer: a buffer containing draw parameters.
//
// offset: the offset to start reading from the draw buffer.
//
// draw_count: the number of draw parameter sets that should be read
// from the draw buffer.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_DrawGPUPrimitivesIndirect
func (pass *GPURenderPass) DrawPrimitivesIndirect(buffer *GPUBuffer, offset int, drawCount int) {
	C.SDL_DrawGPUPrimitivesIndirect((*C.SDL_GPURenderPass)(pass), (*C.SDL_GPUBuffer)(buffer), (C.Uint32)(offset), (C.Uint32)(drawCount))
}

// Draws data using bound graphics state with an index buffer enabled and with
// draw parameters set from a buffer.
//
// The buffer must consist of tightly-packed draw parameter sets that each
// match the layout of SDL_GPUIndexedIndirectDrawCommand. You must not call
// this function before binding a graphics pipeline.
//
// render_pass: a render pass handle.
//
// buffer: a buffer containing draw parameters.
//
// offset: the offset to start reading from the draw buffer.
//
// draw_count: the number of draw parameter sets that should be read
// from the draw buffer.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_DrawGPUIndexedPrimitivesIndirect
func (pass *GPURenderPass) DrawIndexedPrimitivesIndirect(buffer *GPUBuffer, offset int, drawCount int) {
	C.SDL_DrawGPUIndexedPrimitivesIndirect((*C.SDL_GPURenderPass)(pass), (*C.SDL_GPUBuffer)(buffer), (C.Uint32)(offset), (C.Uint32)(drawCount))
}

// Ends the given render pass.
//
// All bound graphics state on the render pass command buffer is unset. The
// render pass handle is now invalid.
//
// render_pass: a render pass handle.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_EndGPURenderPass
func (pass *GPURenderPass) End() {
	C.SDL_EndGPURenderPass((*C.SDL_GPURenderPass)(pass))
}

// Begins a compute pass on a command buffer.
//
// A compute pass is defined by a set of texture subresources and buffers that
// may be written to by compute pipelines. These textures and buffers must
// have been created with the COMPUTE_STORAGE_WRITE bit or the
// COMPUTE_STORAGE_SIMULTANEOUS_READ_WRITE bit. If you do not create a texture
// with COMPUTE_STORAGE_SIMULTANEOUS_READ_WRITE, you must not read from the
// texture in the compute pass. All operations related to compute pipelines
// must take place inside of a compute pass. You must not begin another
// compute pass, or a render pass or copy pass before ending the compute pass.
//
// A VERY IMPORTANT NOTE - Reads and writes in compute passes are NOT
// implicitly synchronized. This means you may cause data races by both
// reading and writing a resource region in a compute pass, or by writing
// multiple times to a resource region. If your compute work depends on
// reading the completed output from a previous dispatch, you MUST end the
// current compute pass and begin a new one before you can safely access the
// data. Otherwise you will receive unexpected results. Reading and writing a
// texture in the same compute pass is only supported by specific texture
// formats. Make sure you check the format support!
//
// command_buffer: a command buffer.
//
// storage_texture_bindings: an array of writeable storage texture
// binding structs.
//
// num_storage_texture_bindings: the number of storage textures to bind
// from the array.
//
// storage_buffer_bindings: an array of writeable storage buffer binding
// structs.
//
// num_storage_buffer_bindings: the number of storage buffers to bind
// from the array.
//
// Returns a compute pass handle.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_BeginGPUComputePass
func (cmd *GPUCommandBuffer) BeginComputePass(storageTextureBindings []*GPUStorageTextureReadWriteBinding, storageBufferBindings []GPUStorageBufferReadWriteBinding) *GPUComputePass {
	textureBindings := make([]C.SDL_GPUStorageTextureReadWriteBinding, len(storageTextureBindings))
	for i, b := range storageTextureBindings {
		textureBindings[i] = C.SDL_GPUStorageTextureReadWriteBinding{
			texture:   (*C.SDL_GPUTexture)(b.Texture),
			mip_level: C.Uint32(b.MipLevel),
			layer:     C.Uint32(b.Layer),
			cycle:     C.bool(b.Cycle),
		}
	}
	bufferBindings := make([]C.SDL_GPUStorageBufferReadWriteBinding, len(storageBufferBindings))
	for i, b := range storageBufferBindings {
		bufferBindings[i] = C.SDL_GPUStorageBufferReadWriteBinding{
			buffer: (*C.SDL_GPUBuffer)(b.Buffer),
			cycle:  C.bool(b.Cycle),
		}
	}
	return (*GPUComputePass)(C.SDL_BeginGPUComputePass((*C.SDL_GPUCommandBuffer)(cmd), unsafe.SliceData(textureBindings), (C.Uint32)(len(textureBindings)), unsafe.SliceData(bufferBindings), (C.Uint32)(len(bufferBindings))))
}

// Binds a compute pipeline on a command buffer for use in compute dispatch.
//
// compute_pass: a compute pass handle.
//
// compute_pipeline: a compute pipeline to bind.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_BindGPUComputePipeline
func (pass *GPUComputePass) BindPipeline(computePipeline *GPUComputePipeline) {
	C.SDL_BindGPUComputePipeline((*C.SDL_GPUComputePass)(pass), (*C.SDL_GPUComputePipeline)(computePipeline))
}

// Binds texture-sampler pairs for use on the compute shader.
//
// The textures must have been created with SDL_GPU_TEXTUREUSAGE_SAMPLER.
//
// compute_pass: a compute pass handle.
//
// first_slot: the compute sampler slot to begin binding from.
//
// texture_sampler_bindings: an array of texture-sampler binding
// structs.
//
// num_bindings: the number of texture-sampler bindings to bind from the
// array.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_BindGPUComputeSamplers
func (pass *GPUComputePass) BindSamplers(firstSlot int, textureSamplerBindings []GPUTextureSamplerBinding) {
	C.SDL_BindGPUComputeSamplers((*C.SDL_GPUComputePass)(pass), (C.Uint32)(firstSlot), (*C.SDL_GPUTextureSamplerBinding)(unsafe.Pointer(unsafe.SliceData(textureSamplerBindings))), (C.Uint32)(len(textureSamplerBindings)))
}

// Binds storage textures as readonly for use on the compute pipeline.
//
// These textures must have been created with
// SDL_GPU_TEXTUREUSAGE_COMPUTE_STORAGE_READ.
//
// compute_pass: a compute pass handle.
//
// first_slot: the compute storage texture slot to begin binding from.
//
// storage_textures: an array of storage textures.
//
// num_bindings: the number of storage textures to bind from the array.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_BindGPUComputeStorageTextures
func (pass *GPUComputePass) BindStorageTextures(firstSlot int, storageTextures []*GPUTexture) {
	C.SDL_BindGPUComputeStorageTextures((*C.SDL_GPUComputePass)(pass), (C.Uint32)(firstSlot), (**C.SDL_GPUTexture)(unsafe.Pointer(unsafe.SliceData(storageTextures))), (C.Uint32)(len(storageTextures)))
}

// Binds storage buffers as readonly for use on the compute pipeline.
//
// These buffers must have been created with
// SDL_GPU_BUFFERUSAGE_COMPUTE_STORAGE_READ.
//
// compute_pass: a compute pass handle.
//
// first_slot: the compute storage buffer slot to begin binding from.
//
// storage_buffers: an array of storage buffer binding structs.
//
// num_bindings: the number of storage buffers to bind from the array.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_BindGPUComputeStorageBuffers
func (pass *GPUComputePass) BindStorageBuffers(firstSlot int, storageBuffers []*GPUBuffer) {
	C.SDL_BindGPUComputeStorageBuffers((*C.SDL_GPUComputePass)(pass), (C.Uint32)(firstSlot), (**C.SDL_GPUBuffer)(unsafe.Pointer(unsafe.SliceData(storageBuffers))), (C.Uint32)(len(storageBuffers)))
}

// Dispatches compute work.
//
// You must not call this function before binding a compute pipeline.
//
// A VERY IMPORTANT NOTE If you dispatch multiple times in a compute pass, and
// the dispatches write to the same resource region as each other, there is no
// guarantee of which order the writes will occur. If the write order matters,
// you MUST end the compute pass and begin another one.
//
// compute_pass: a compute pass handle.
//
// groupcount_x: number of local workgroups to dispatch in the X
// dimension.
//
// groupcount_y: number of local workgroups to dispatch in the Y
// dimension.
//
// groupcount_z: number of local workgroups to dispatch in the Z
// dimension.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_DispatchGPUCompute
func (pass *GPUComputePass) Dispatch(groupcountX int, groupcountY int, groupcountZ int) {
	C.SDL_DispatchGPUCompute((*C.SDL_GPUComputePass)(pass), (C.Uint32)(groupcountX), (C.Uint32)(groupcountY), (C.Uint32)(groupcountZ))
}

// Dispatches compute work with parameters set from a buffer.
//
// The buffer layout should match the layout of
// SDL_GPUIndirectDispatchCommand. You must not call this function before
// binding a compute pipeline.
//
// A VERY IMPORTANT NOTE If you dispatch multiple times in a compute pass, and
// the dispatches write to the same resource region as each other, there is no
// guarantee of which order the writes will occur. If the write order matters,
// you MUST end the compute pass and begin another one.
//
// compute_pass: a compute pass handle.
//
// buffer: a buffer containing dispatch parameters.
//
// offset: the offset to start reading from the dispatch buffer.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_DispatchGPUComputeIndirect
func (pass *GPUComputePass) DispatchIndirect(buffer *GPUBuffer, offset uint32) {
	C.SDL_DispatchGPUComputeIndirect((*C.SDL_GPUComputePass)(pass), (*C.SDL_GPUBuffer)(buffer), (C.Uint32)(offset))
}

// Ends the current compute pass.
//
// All bound compute state on the command buffer is unset. The compute pass
// handle is now invalid.
//
// compute_pass: a compute pass handle.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_EndGPUComputePass
func (pass *GPUComputePass) End() {
	C.SDL_EndGPUComputePass((*C.SDL_GPUComputePass)(pass))
}

// Maps a transfer buffer into application address space.
//
// You must unmap the transfer buffer before encoding upload commands.
//
// device: a GPU context.
//
// transfer_buffer: a transfer buffer.
//
// cycle: if true, cycles the transfer buffer if it is already bound.
//
// Returns the address of the mapped transfer buffer memory, or NULL on
// failure; call SDL_GetError() for more information.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_MapGPUTransferBuffer
func (device *GPUDevice) MapTransferBuffer(transferBuffer *GPUTransferBuffer, cycle bool) (*GPUTransferBufferMapping, error) {
	if transferBuffer.mapping != nil {
		transferBuffer.mapping.p = nil
		transferBuffer.mapping = nil
	}
	p := C.SDL_MapGPUTransferBuffer((*C.SDL_GPUDevice)(device), transferBuffer.internal, (C.bool)(cycle))
	if p == nil {
		return nil, getError()
	}
	m := &GPUTransferBufferMapping{device, transferBuffer, p}
	transferBuffer.mapping = m
	return m, nil
}

// GPUTransferBufferMapping represents an region of memory mapped from a
// transfer buffer.
type GPUTransferBufferMapping struct {
	device         *GPUDevice
	transferBuffer *GPUTransferBuffer
	p              unsafe.Pointer
}

// Mapped reports wether the mapping is valid.
func (m *GPUTransferBufferMapping) Mapped() bool {
	return m.p != nil
}

// WriteAt writes data into the mapped memory region at the given offset.
func (m *GPUTransferBufferMapping) WriteAt(offset int, data []byte) {
	if m.p == nil {
		return
	}
	copy(unsafe.Slice((*byte)(m.p), m.transferBuffer.size)[offset:], data)
}

// ReadAt reads data from the mapped memory region at the given offset.
func (m *GPUTransferBufferMapping) ReadAt(offset int, data []byte) {
	if m.p == nil {
		return
	}
	copy(data, unsafe.Slice((*byte)(m.p), m.transferBuffer.size)[offset:])
}

// Size returns the size of the mapped memory in bytes.
func (m *GPUTransferBufferMapping) Size() int {
	return m.transferBuffer.size
}

// TransferBuffer returns the [GPUTransferBuffer] from which this mapping was
// created.
func (m *GPUTransferBufferMapping) TransferBuffer() *GPUTransferBuffer {
	return m.transferBuffer
}

// Unmaps a previously mapped transfer buffer.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_UnmapGPUTransferBuffer
func (m *GPUTransferBufferMapping) Unmap() {
	C.SDL_UnmapGPUTransferBuffer((*C.SDL_GPUDevice)(m.device), m.transferBuffer.internal)
	m.p = nil
	m.transferBuffer.mapping = nil
}

// Begins a copy pass on a command buffer.
//
// All operations related to copying to or from buffers or textures take place
// inside a copy pass. You must not begin another copy pass, or a render pass
// or compute pass before ending the copy pass.
//
// command_buffer: a command buffer.
//
// Returns a copy pass handle.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_BeginGPUCopyPass
func (cmd *GPUCommandBuffer) BeginCopyPass() *GPUCopyPass {
	return (*GPUCopyPass)(C.SDL_BeginGPUCopyPass((*C.SDL_GPUCommandBuffer)(cmd)))
}

// Uploads data from a transfer buffer to a texture.
//
// The upload occurs on the GPU timeline. You may assume that the upload has
// finished in subsequent commands.
//
// You must align the data in the transfer buffer to a multiple of the texel
// size of the texture format.
//
// copy_pass: a copy pass handle.
//
// source: the source transfer buffer with image layout information.
//
// destination: the destination texture region.
//
// cycle: if true, cycles the texture if the texture is bound, otherwise
// overwrites the data.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_UploadToGPUTexture
func (pass *GPUCopyPass) UploadToTexture(source *GPUTextureTransferInfo, destination *GPUTextureRegion, cycle bool) {
	C.SDL_UploadToGPUTexture((*C.SDL_GPUCopyPass)(pass), &C.SDL_GPUTextureTransferInfo{
		transfer_buffer: source.TransferBuffer.internal,
		offset:          C.Uint32(source.Offset),
		pixels_per_row:  C.Uint32(source.PixelsPerRow),
		rows_per_layer:  C.Uint32(source.RowsPerLayer),
	}, &C.SDL_GPUTextureRegion{
		texture:   (*C.SDL_GPUTexture)(destination.Texture),
		mip_level: C.Uint32(destination.MipLevel),
		layer:     C.Uint32(destination.Layer),
		x:         C.Uint32(destination.X),
		y:         C.Uint32(destination.Y),
		z:         C.Uint32(destination.Z),
		w:         C.Uint32(destination.W),
		h:         C.Uint32(destination.H),
		d:         C.Uint32(destination.D),
	}, (C.bool)(cycle))
}

// Uploads data from a transfer buffer to a buffer.
//
// The upload occurs on the GPU timeline. You may assume that the upload has
// finished in subsequent commands.
//
// copy_pass: a copy pass handle.
//
// source: the source transfer buffer with offset.
//
// destination: the destination buffer with offset and size.
//
// cycle: if true, cycles the buffer if it is already bound, otherwise
// overwrites the data.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_UploadToGPUBuffer
func (pass *GPUCopyPass) UploadToBuffer(source *GPUTransferBufferLocation, destination *GPUBufferRegion, cycle bool) {
	C.SDL_UploadToGPUBuffer((*C.SDL_GPUCopyPass)(pass), &C.SDL_GPUTransferBufferLocation{
		transfer_buffer: source.TransferBuffer.internal,
		offset:          C.Uint32(source.Offset),
	}, &C.SDL_GPUBufferRegion{
		buffer: (*C.SDL_GPUBuffer)(destination.Buffer),
		offset: C.Uint32(destination.Offset),
		size:   C.Uint32(destination.Size),
	}, (C.bool)(cycle))
}

// Performs a texture-to-texture copy.
//
// This copy occurs on the GPU timeline. You may assume the copy has finished
// in subsequent commands.
//
// copy_pass: a copy pass handle.
//
// source: a source texture region.
//
// destination: a destination texture region.
//
// w: the width of the region to copy.
//
// h: the height of the region to copy.
//
// d: the depth of the region to copy.
//
// cycle: if true, cycles the destination texture if the destination
// texture is bound, otherwise overwrites the data.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_CopyGPUTextureToTexture
func (pass *GPUCopyPass) CopyTextureToTexture(source *GPUTextureLocation, destination *GPUTextureLocation, w int, h int, d int, cycle bool) {
	C.SDL_CopyGPUTextureToTexture((*C.SDL_GPUCopyPass)(pass), &C.SDL_GPUTextureLocation{
		texture:   (*C.SDL_GPUTexture)(source.Texture),
		mip_level: C.Uint32(source.MipLevel),
		layer:     C.Uint32(source.Layer),
		x:         C.Uint32(source.X),
		y:         C.Uint32(source.Y),
		z:         C.Uint32(source.Z),
	}, &C.SDL_GPUTextureLocation{
		texture:   (*C.SDL_GPUTexture)(destination.Texture),
		mip_level: C.Uint32(destination.MipLevel),
		layer:     C.Uint32(destination.Layer),
		x:         C.Uint32(destination.X),
		y:         C.Uint32(destination.Y),
		z:         C.Uint32(destination.Z),
	}, (C.Uint32)(w), (C.Uint32)(h), (C.Uint32)(d), (C.bool)(cycle))
}

// Performs a buffer-to-buffer copy.
//
// This copy occurs on the GPU timeline. You may assume the copy has finished
// in subsequent commands.
//
// copy_pass: a copy pass handle.
//
// source: the buffer and offset to copy from.
//
// destination: the buffer and offset to copy to.
//
// size: the length of the buffer to copy.
//
// cycle: if true, cycles the destination buffer if it is already bound,
// otherwise overwrites the data.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_CopyGPUBufferToBuffer
func (pass *GPUCopyPass) CopyBufferToBuffer(source *GPUBufferLocation, destination *GPUBufferLocation, size int, cycle bool) {
	C.SDL_CopyGPUBufferToBuffer((*C.SDL_GPUCopyPass)(pass), &C.SDL_GPUBufferLocation{
		buffer: (*C.SDL_GPUBuffer)(source.Buffer),
		offset: C.Uint32(source.Offset),
	}, &C.SDL_GPUBufferLocation{
		buffer: (*C.SDL_GPUBuffer)(destination.Buffer),
		offset: C.Uint32(destination.Offset),
	}, (C.Uint32)(size), (C.bool)(cycle))
}

// Copies data from a texture to a transfer buffer on the GPU timeline.
//
// This data is not guaranteed to be copied until the command buffer fence is
// signaled.
//
// copy_pass: a copy pass handle.
//
// source: the source texture region.
//
// destination: the destination transfer buffer with image layout
// information.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_DownloadFromGPUTexture
func (pass *GPUCopyPass) DownloadFromTexture(source *GPUTextureRegion, destination *GPUTextureTransferInfo) {
	C.SDL_DownloadFromGPUTexture((*C.SDL_GPUCopyPass)(pass), &C.SDL_GPUTextureRegion{
		texture:   (*C.SDL_GPUTexture)(source.Texture),
		mip_level: C.Uint32(source.MipLevel),
		layer:     C.Uint32(source.Layer),
		x:         C.Uint32(source.X),
		y:         C.Uint32(source.Y),
		z:         C.Uint32(source.Z),
		w:         C.Uint32(source.W),
		h:         C.Uint32(source.H),
		d:         C.Uint32(source.D),
	}, &C.SDL_GPUTextureTransferInfo{
		transfer_buffer: destination.TransferBuffer.internal,
		offset:          C.Uint32(destination.Offset),
		pixels_per_row:  C.Uint32(destination.PixelsPerRow),
		rows_per_layer:  C.Uint32(destination.RowsPerLayer),
	})
}

// Copies data from a buffer to a transfer buffer on the GPU timeline.
//
// This data is not guaranteed to be copied until the command buffer fence is
// signaled.
//
// copy_pass: a copy pass handle.
//
// source: the source buffer with offset and size.
//
// destination: the destination transfer buffer with offset.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_DownloadFromGPUBuffer
func (pass *GPUCopyPass) DownloadFromBuffer(source *GPUBufferRegion, destination *GPUTransferBufferLocation) {
	C.SDL_DownloadFromGPUBuffer((*C.SDL_GPUCopyPass)(pass), &C.SDL_GPUBufferRegion{
		buffer: (*C.SDL_GPUBuffer)(source.Buffer),
		offset: C.Uint32(source.Offset),
		size:   C.Uint32(source.Size),
	}, &C.SDL_GPUTransferBufferLocation{
		transfer_buffer: destination.TransferBuffer.internal,
		offset:          C.Uint32(destination.Offset),
	})
}

// Ends the current copy pass.
//
// copy_pass: a copy pass handle.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_EndGPUCopyPass
func (pass *GPUCopyPass) End() {
	C.SDL_EndGPUCopyPass((*C.SDL_GPUCopyPass)(pass))
}

// Generates mipmaps for the given texture.
//
// This function must not be called inside of any pass.
//
// command_buffer: a command_buffer.
//
// texture: a texture with more than 1 mip level.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GenerateMipmapsForGPUTexture
func (cmd *GPUCommandBuffer) GenerateMipmapsForTexture(texture *GPUTexture) {
	C.SDL_GenerateMipmapsForGPUTexture((*C.SDL_GPUCommandBuffer)(cmd), (*C.SDL_GPUTexture)(texture))
}

// Blits from a source texture region to a destination texture region.
//
// This function must not be called inside of any pass.
//
// command_buffer: a command buffer.
//
// info: the blit info struct containing the blit parameters.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_BlitGPUTexture
func (cmd *GPUCommandBuffer) BlitTexture(info *GPUBlitInfo) {
	C.SDL_BlitGPUTexture((*C.SDL_GPUCommandBuffer)(cmd), &C.SDL_GPUBlitInfo{
		source: C.SDL_GPUBlitRegion{
			texture:              (*C.SDL_GPUTexture)(info.Source.Texture),
			mip_level:            C.Uint32(info.Source.MipLevel),
			layer_or_depth_plane: C.Uint32(info.Source.LayerOrDepthPlane),
			x:                    C.Uint32(info.Source.X),
			y:                    C.Uint32(info.Source.Y),
			w:                    C.Uint32(info.Source.W),
			h:                    C.Uint32(info.Source.H),
		},
		destination: C.SDL_GPUBlitRegion{
			texture:              (*C.SDL_GPUTexture)(info.Destination.Texture),
			mip_level:            C.Uint32(info.Destination.MipLevel),
			layer_or_depth_plane: C.Uint32(info.Destination.LayerOrDepthPlane),
			x:                    C.Uint32(info.Destination.X),
			y:                    C.Uint32(info.Destination.Y),
			w:                    C.Uint32(info.Destination.W),
			h:                    C.Uint32(info.Destination.H),
		},
		load_op:     C.SDL_GPULoadOp(info.LoadOp),
		clear_color: C.SDL_FColor{C.float(info.ClearColor.R), C.float(info.ClearColor.G), C.float(info.ClearColor.B), C.float(info.ClearColor.A)},
		flip_mode:   C.SDL_FlipMode(info.FlipMode),
		filter:      C.SDL_GPUFilter(info.Filter),
		cycle:       C.bool(info.Cycle),
	})
}

// Determines whether a swapchain composition is supported by the window.
//
// The window must be claimed before calling this function.
//
// device: a GPU context.
//
// window: an SDL_Window.
//
// swapchain_composition: the swapchain composition to check.
//
// Returns true if supported, false if unsupported.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_WindowSupportsGPUSwapchainComposition
func (device *GPUDevice) WindowSupportsSwapchainComposition(window *Window, swapchainComposition GPUSwapchainComposition) bool {
	return (bool)(C.SDL_WindowSupportsGPUSwapchainComposition((*C.SDL_GPUDevice)(device), (*C.SDL_Window)(window), (C.SDL_GPUSwapchainComposition)(swapchainComposition)))
}

// Determines whether a presentation mode is supported by the window.
//
// The window must be claimed before calling this function.
//
// device: a GPU context.
//
// window: an SDL_Window.
//
// present_mode: the presentation mode to check.
//
// Returns true if supported, false if unsupported.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_WindowSupportsGPUPresentMode
func (device *GPUDevice) WindowSupportsPresentMode(window *Window, presentMode GPUPresentMode) bool {
	return (bool)(C.SDL_WindowSupportsGPUPresentMode((*C.SDL_GPUDevice)(device), (*C.SDL_Window)(window), (C.SDL_GPUPresentMode)(presentMode)))
}

// Claims a window, creating a swapchain structure for it.
//
// This must be called before SDL_AcquireGPUSwapchainTexture is called using
// the window. You should only call this function from the thread that created
// the window.
//
// The swapchain will be created with SDL_GPU_SWAPCHAINCOMPOSITION_SDR and
// SDL_GPU_PRESENTMODE_VSYNC. If you want to have different swapchain
// parameters, you must call SDL_SetGPUSwapchainParameters after claiming the
// window.
//
// device: a GPU context.
//
// window: an SDL_Window.
//
// Returns true on success, or false on failure; call SDL_GetError() for more
// information.
//
// This function should only be called from the thread that
// created the window.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_ClaimWindowForGPUDevice
func (device *GPUDevice) ClaimWindow(window *Window) error {
	if !C.SDL_ClaimWindowForGPUDevice((*C.SDL_GPUDevice)(device), (*C.SDL_Window)(window)) {
		return getError()
	}
	return nil
}

// Unclaims a window, destroying its swapchain structure.
//
// device: a GPU context.
//
// window: an SDL_Window that has been claimed.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_ReleaseWindowFromGPUDevice
func (device *GPUDevice) ReleaseWindowFromGPUDevice(window *Window) {
	C.SDL_ReleaseWindowFromGPUDevice((*C.SDL_GPUDevice)(device), (*C.SDL_Window)(window))
}

// Changes the swapchain parameters for the given claimed window.
//
// This function will fail if the requested present mode or swapchain
// composition are unsupported by the device. Check if the parameters are
// supported via SDL_WindowSupportsGPUPresentMode /
// SDL_WindowSupportsGPUSwapchainComposition prior to calling this function.
//
// SDL_GPU_PRESENTMODE_VSYNC and SDL_GPU_SWAPCHAINCOMPOSITION_SDR are always
// supported.
//
// device: a GPU context.
//
// window: an SDL_Window that has been claimed.
//
// swapchain_composition: the desired composition of the swapchain.
//
// present_mode: the desired present mode for the swapchain.
//
// Returns true if successful, false on error; call SDL_GetError() for more
// information.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetGPUSwapchainParameters
func (device *GPUDevice) SetSwapchainParameters(window *Window, swapchainComposition GPUSwapchainComposition, presentMode GPUPresentMode) error {
	if !C.SDL_SetGPUSwapchainParameters((*C.SDL_GPUDevice)(device), (*C.SDL_Window)(window), (C.SDL_GPUSwapchainComposition)(swapchainComposition), (C.SDL_GPUPresentMode)(presentMode)) {
		return getError()
	}
	return nil
}

// Configures the maximum allowed number of frames in flight.
//
// The default value when the device is created is 2. This means that after
// you have submitted 2 frames for presentation, if the GPU has not finished
// working on the first frame, SDL_AcquireGPUSwapchainTexture() will fill the
// swapchain texture pointer with NULL, and
// SDL_WaitAndAcquireGPUSwapchainTexture() will block.
//
// Higher values increase throughput at the expense of visual latency. Lower
// values decrease visual latency at the expense of throughput.
//
// Note that calling this function will stall and flush the command queue to
// prevent synchronization issues.
//
// The minimum value of allowed frames in flight is 1, and the maximum is 3.
//
// device: a GPU context.
//
// allowed_frames_in_flight: the maximum number of frames that can be
// pending on the GPU.
//
// Returns true if successful, false on error; call SDL_GetError() for more
// information.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetGPUAllowedFramesInFlight
func (device *GPUDevice) SetAllowedFramesInFlight(allowedFramesInFlight int) error {
	if !C.SDL_SetGPUAllowedFramesInFlight((*C.SDL_GPUDevice)(device), (C.Uint32)(allowedFramesInFlight)) {
		return getError()
	}
	return nil
}

// Obtains the texture format of the swapchain for the given window.
//
// Note that this format can change if the swapchain parameters change.
//
// device: a GPU context.
//
// window: an SDL_Window that has been claimed.
//
// Returns the texture format of the swapchain.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetGPUSwapchainTextureFormat
func (device *GPUDevice) GetSwapchainTextureFormat(window *Window) GPUTextureFormat {
	return (GPUTextureFormat)(C.SDL_GetGPUSwapchainTextureFormat((*C.SDL_GPUDevice)(device), (*C.SDL_Window)(window)))
}

// Acquire a texture to use in presentation.
//
// When a swapchain texture is acquired on a command buffer, it will
// automatically be submitted for presentation when the command buffer is
// submitted. The swapchain texture should only be referenced by the command
// buffer used to acquire it.
//
// This function will fill the swapchain texture handle with NULL if too many
// frames are in flight. This is not an error.
//
// If you use this function, it is possible to create a situation where many
// command buffers are allocated while the rendering context waits for the GPU
// to catch up, which will cause memory usage to grow. You should use
// SDL_WaitAndAcquireGPUSwapchainTexture() unless you know what you are doing
// with timing.
//
// The swapchain texture is managed by the implementation and must not be
// freed by the user. You MUST NOT call this function from any thread other
// than the one that created the window.
//
// command_buffer: a command buffer.
//
// window: a window that has been claimed.
//
// swapchain_texture: a pointer filled in with a swapchain texture
// handle.
//
// swapchain_texture_width: a pointer filled in with the swapchain
// texture width, may be NULL.
//
// swapchain_texture_height: a pointer filled in with the swapchain
// texture height, may be NULL.
//
// Returns true on success, false on error; call SDL_GetError() for more
// information.
//
// This function should only be called from the thread that
// created the window.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_AcquireGPUSwapchainTexture
func (cmd *GPUCommandBuffer) AcquireSwapchainTexture(window *Window) (swapchainTexture *GPUTexture, width, height int, err error) {
	var texture *C.SDL_GPUTexture
	var w, h C.Uint32
	ok := C.SDL_AcquireGPUSwapchainTexture((*C.SDL_GPUCommandBuffer)(cmd), (*C.SDL_Window)(window), &texture, &w, &h)
	if !ok {
		return nil, 0, 0, getError()
	}
	return (*GPUTexture)(texture), int(w), int(h), nil
}

// Blocks the thread until a swapchain texture is available to be acquired.
//
// device: a GPU context.
//
// window: a window that has been claimed.
//
// Returns true on success, false on failure; call SDL_GetError() for more
// information.
//
// This function should only be called from the thread that
// created the window.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_WaitForGPUSwapchain
func (device *GPUDevice) WaitForGPUSwapchain(window *Window) bool {
	return (bool)(C.SDL_WaitForGPUSwapchain((*C.SDL_GPUDevice)(device), (*C.SDL_Window)(window)))
}

// Blocks the thread until a swapchain texture is available to be acquired,
// and then acquires it.
//
// When a swapchain texture is acquired on a command buffer, it will
// automatically be submitted for presentation when the command buffer is
// submitted. The swapchain texture should only be referenced by the command
// buffer used to acquire it. It is an error to call
// SDL_CancelGPUCommandBuffer() after a swapchain texture is acquired.
//
// This function can fill the swapchain texture handle with NULL in certain
// cases, for example if the window is minimized. This is not an error. You
// should always make sure to check whether the pointer is NULL before
// actually using it.
//
// The swapchain texture is managed by the implementation and must not be
// freed by the user. You MUST NOT call this function from any thread other
// than the one that created the window.
//
// command_buffer: a command buffer.
//
// window: a window that has been claimed.
//
// swapchain_texture: a pointer filled in with a swapchain texture
// handle.
//
// swapchain_texture_width: a pointer filled in with the swapchain
// texture width, may be NULL.
//
// swapchain_texture_height: a pointer filled in with the swapchain
// texture height, may be NULL.
//
// Returns true on success, false on error; call SDL_GetError() for more
// information.
//
// This function should only be called from the thread that
// created the window.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_WaitAndAcquireGPUSwapchainTexture
func (cmd *GPUCommandBuffer) WaitAndAcquireSwapchainTexture(window *Window) (swapchainTexture *GPUTexture, width, height int, err error) {
	var texture *C.SDL_GPUTexture
	var w, h C.Uint32
	ok := C.SDL_WaitAndAcquireGPUSwapchainTexture((*C.SDL_GPUCommandBuffer)(cmd), (*C.SDL_Window)(window), &texture, &w, &h)
	if !ok {
		return nil, 0, 0, getError()
	}
	return (*GPUTexture)(texture), int(w), int(h), nil
}

// Submits a command buffer so its commands can be processed on the GPU.
//
// It is invalid to use the command buffer after this is called.
//
// This must be called from the thread the command buffer was acquired on.
//
// All commands in the submission are guaranteed to begin executing before any
// command in a subsequent submission begins executing.
//
// command_buffer: a command buffer.
//
// Returns true on success, false on failure; call SDL_GetError() for more
// information.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SubmitGPUCommandBuffer
func (cmd *GPUCommandBuffer) Submit() error {
	if !C.SDL_SubmitGPUCommandBuffer((*C.SDL_GPUCommandBuffer)(cmd)) {
		return getError()
	}
	return nil
}

// Submits a command buffer so its commands can be processed on the GPU, and
// acquires a fence associated with the command buffer.
//
// You must release this fence when it is no longer needed or it will cause a
// leak. It is invalid to use the command buffer after this is called.
//
// This must be called from the thread the command buffer was acquired on.
//
// All commands in the submission are guaranteed to begin executing before any
// command in a subsequent submission begins executing.
//
// command_buffer: a command buffer.
//
// Returns a fence associated with the command buffer, or NULL on failure;
// call SDL_GetError() for more information.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SubmitGPUCommandBufferAndAcquireFence
func (cmd *GPUCommandBuffer) SubmitAndAcquireFence() (*GPUFence, error) {
	fence := (*GPUFence)(C.SDL_SubmitGPUCommandBufferAndAcquireFence((*C.SDL_GPUCommandBuffer)(cmd)))
	if fence == nil {
		return nil, getError()
	}
	return fence, nil
}

// Cancels a command buffer.
//
// None of the enqueued commands are executed.
//
// It is an error to call this function after a swapchain texture has been
// acquired.
//
// This must be called from the thread the command buffer was acquired on.
//
// You must not reference the command buffer after calling this function.
//
// command_buffer: a command buffer.
//
// Returns true on success, false on error; call SDL_GetError() for more
// information.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_CancelGPUCommandBuffer
func (cmd *GPUCommandBuffer) Cancel() error {
	if !C.SDL_CancelGPUCommandBuffer((*C.SDL_GPUCommandBuffer)(cmd)) {
		return getError()
	}
	return nil
}

// Blocks the thread until the GPU is completely idle.
//
// device: a GPU context.
//
// Returns true on success, false on failure; call SDL_GetError() for more
// information.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_WaitForGPUIdle
func (device *GPUDevice) WaitForIdle() error {
	if !C.SDL_WaitForGPUIdle((*C.SDL_GPUDevice)(device)) {
		return getError()
	}
	return nil
}

// Blocks the thread until the given fences are signaled.
//
// device: a GPU context.
//
// wait_all: if 0, wait for any fence to be signaled, if 1, wait for all
// fences to be signaled.
//
// fences: an array of fences to wait on.
//
// num_fences: the number of fences in the fences array.
//
// Returns true on success, false on failure; call SDL_GetError() for more
// information.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_WaitForGPUFences
func (device *GPUDevice) WaitForFences(waitAll bool, fences []*GPUFence) error {
	if !C.SDL_WaitForGPUFences((*C.SDL_GPUDevice)(device), (C.bool)(waitAll), (**C.SDL_GPUFence)(unsafe.Pointer(unsafe.SliceData(fences))), (C.Uint32)(len(fences))) {
		return getError()
	}
	return nil
}

// Checks the status of a fence.
//
// device: a GPU context.
//
// fence: a fence.
//
// Returns true if the fence is signaled, false if it is not.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_QueryGPUFence
func (device *GPUDevice) QueryFence(fence *GPUFence) bool {
	return (bool)(C.SDL_QueryGPUFence((*C.SDL_GPUDevice)(device), (*C.SDL_GPUFence)(fence)))
}

// Releases a fence obtained from SDL_SubmitGPUCommandBufferAndAcquireFence.
//
// device: a GPU context.
//
// fence: a fence.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_ReleaseGPUFence
func (device *GPUDevice) ReleaseFence(fence *GPUFence) {
	C.SDL_ReleaseGPUFence((*C.SDL_GPUDevice)(device), (*C.SDL_GPUFence)(fence))
}

// Obtains the texel block size for a texture format.
//
// format: the texture format you want to know the texel size of.
//
// Returns the texel block size of the texture format.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GPUTextureFormatTexelBlockSize
func GPUTextureFormatTexelBlockSize(format GPUTextureFormat) int {
	return (int)(C.SDL_GPUTextureFormatTexelBlockSize((C.SDL_GPUTextureFormat)(format)))
}

// Determines whether a texture format is supported for a given type and
// usage.
//
// device: a GPU context.
//
// format: the texture format to check.
//
// type: the type of texture (2D, 3D, Cube).
//
// usage: a bitmask of all usage scenarios to check.
//
// Returns whether the texture format is supported for this type and usage.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GPUTextureSupportsFormat
func (device *GPUDevice) TextureSupportsFormat(format GPUTextureFormat, typ GPUTextureType, usage GPUTextureUsageFlags) bool {
	return (bool)(C.SDL_GPUTextureSupportsFormat((*C.SDL_GPUDevice)(device), (C.SDL_GPUTextureFormat)(format), (C.SDL_GPUTextureType)(typ), (C.SDL_GPUTextureUsageFlags)(usage)))
}

// Determines if a sample count for a texture format is supported.
//
// device: a GPU context.
//
// format: the texture format to check.
//
// sample_count: the sample count to check.
//
// Returns a hardware-specific version of min(preferred, possible).
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GPUTextureSupportsSampleCount
func (device *GPUDevice) GPUTextureSupportsSampleCount(format GPUTextureFormat, sampleCount GPUSampleCount) bool {
	return (bool)(C.SDL_GPUTextureSupportsSampleCount((*C.SDL_GPUDevice)(device), (C.SDL_GPUTextureFormat)(format), (C.SDL_GPUSampleCount)(sampleCount)))
}

// Calculate the size in bytes of a texture format with dimensions.
//
// format: a texture format.
//
// width: width in pixels.
//
// height: height in pixels.
//
// depth_or_layer_count: depth for 3D textures or layer count otherwise.
//
// Returns the size of a texture with this format and dimensions.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_CalculateGPUTextureFormatSize
func CalculateGPUTextureFormatSize(format GPUTextureFormat, width int, height int, depthOrLayerCount int) int {
	return (int)(C.SDL_CalculateGPUTextureFormatSize((C.SDL_GPUTextureFormat)(format), (C.Uint32)(width), (C.Uint32)(height), (C.Uint32)(depthOrLayerCount)))
}
