
package fields

/*

2020-09-16
the exif package does not currently work on video formats

*/

var MatchNames = []string{ 
	"*.jpg",
	"*.JPG",
	"*.jpeg",
	"*.JPEG",
	"*.png",
	"*.PNG",
	//"*.mp3",
	//"*.MP3",
	//"*.mp4",
	//"*.MP4",
	//"*.wmv",
	//"*.WMV",
	//"*.mov",
	//"*.MOV",
	"*.tiff",
	"*.TIFF",
}

var ExifFields = []string {
	"Make",
	"Model",
	"Software",
	"Date",
	"Desc",
	"Orientation",
	"ExposureTime",
	"FNumber",
	"ShutterSpeedValue",
	"ApertureValue",
	"ExposureProgram",
	"SpectralSensitivity",
	"ISOSpeedRatings",
	"FocalLength",
	"DigitalZoomRatio",
	"Flash",
	"OECF",
	"BrightnessValue",
	"ExposureBiasValue",
	"MaxApertureValue",
	"SubjectDistance",
	"MeteringMode",
	"LightSource",
	"ExposureMode",
	"WhiteBalance",
	"Contrast",
	"SubjectDistanceRange",
	"ExposureBiasValue",
	"PixelXDimension",
	"PixelYDimension",
	"XResolution",
	"YResolution",
	"CompressedBitsPerPixel",
}

