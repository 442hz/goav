// Created by cgo -godefs - DO NOT EDIT
// cgo -godefs frame.go

package avutil

// import (
// 	"unsafe"
// )

type (
	// AvBuffer	struct{}
	// AvBufferRef	struct {
	// 	Buffer		*AvBuffer
	// 	Data		*uint8
	// 	Size		int32
	// 	Pad_cgo_0	[4]byte
	// }
	// AvBufferPool	struct{}
	Frame struct {
		Data                   [8]*uint8
		Linesize               [8]int32
		Extended_data          **uint8
		Width                  int32
		Height                 int32
		Nb_samples             int32
		Format                 int32
		Key_frame              int32
		Pict_type              uint32
		Sample_aspect_ratio    _Ctype_struct_AVRational
		Pts                    int64
		Pkt_pts                int64
		Pkt_dts                int64
		Coded_picture_number   int32
		Display_picture_number int32
		Quality                int32
		Pad_cgo_0              [4]byte
		Opaque                 *byte
		Error                  [8]uint64
		Repeat_pict            int32
		Interlaced_frame       int32
		Top_field_first        int32
		Palette_has_changed    int32
		Reordered_opaque       int64
		Sample_rate            int32
		Pad_cgo_1              [4]byte
		Channel_layout         uint64
		Buf                    [8]*AvBufferRef
		Extended_buf           **AvBufferRef
		Nb_extended_buf        int32
		Pad_cgo_2              [4]byte
		Side_data              **AvFrameSideData
		Nb_side_data           int32
		Flags                  int32
		Color_range            uint32
		Color_primaries        uint32
		Color_trc              uint32
		Colorspace             uint32
		Chroma_location        uint32
		Pad_cgo_3              [4]byte
		Best_effort_timestamp  int64
		Pkt_pos                int64
		Pkt_duration           int64
		Metadata               *_Ctype_struct_AVDictionary
		Decode_error_flags     int32
		Channels               int32
		Pkt_size               int32
		Pad_cgo_4              [4]byte
		Qscale_table           *int8
		Qstride                int32
		Qscale_type            int32
		Qp_table_buf           *AvBufferRef
		Hw_frames_ctx          *AvBufferRef
	}
	// AvFrameSideData	struct {
	// 	Type		uint32
	// 	Pad_cgo_0	[4]byte
	// 	Data		*uint8
	// 	Size		int32
	// 	Pad_cgo_1	[4]byte
	// 	Metadata	*_Ctype_struct_AVDictionary
	// 	Buf		*AvBufferRef
	// }
	// AvFrameSideDataType	uint32
)
