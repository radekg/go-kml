package kml

import (
	"image/color"
	"time"
)

// An AltitudeModeEnum is an altitudeModeEnumType.
type AltitudeModeEnum string

// AltitudeModeEnums.
const (
	AltitudeModeClampToGround    AltitudeModeEnum = "clampToGround"
	AltitudeModeRelativeToGround AltitudeModeEnum = "relativeToGround"
	AltitudeModeAbsolute         AltitudeModeEnum = "absolute"
)

// A ColorModeEnum is a colorModeEnumType.
type ColorModeEnum string

// ColorModeEnums.
const (
	ColorModeNormal ColorModeEnum = "normal"
	ColorModeRandom ColorModeEnum = "random"
)

// A DisplayModeEnum is a displayModeEnumType.
type DisplayModeEnum string

// DisplayModeEnums.
const (
	DisplayModeDefault DisplayModeEnum = "default"
	DisplayModeHide    DisplayModeEnum = "hide"
)

// A GridOriginEnum is a gridOriginEnumType.
type GridOriginEnum string

// GridOriginEnums.
const (
	GridOriginLowerLeft GridOriginEnum = "lowerLeft"
	GridOriginUpperLeft GridOriginEnum = "upperLeft"
)

// A ItemIconStateEnum is a itemIconStateEnumType.
type ItemIconStateEnum string

// ItemIconStateEnums.
const (
	ItemIconStateOpen      ItemIconStateEnum = "open"
	ItemIconStateClosed    ItemIconStateEnum = "closed"
	ItemIconStateError     ItemIconStateEnum = "error"
	ItemIconStateFetching0 ItemIconStateEnum = "fetching0"
	ItemIconStateFetching1 ItemIconStateEnum = "fetching1"
	ItemIconStateFetching2 ItemIconStateEnum = "fetching2"
)

// A ListItemTypeEnum is a listItemTypeEnumType.
type ListItemTypeEnum string

// ListItemTypeEnums.
const (
	ListItemTypeRadioFolder       ListItemTypeEnum = "radioFolder"
	ListItemTypeCheck             ListItemTypeEnum = "check"
	ListItemTypeCheckHideChildren ListItemTypeEnum = "checkHideChildren"
	ListItemTypeCheckOffOnly      ListItemTypeEnum = "checkOffOnly"
)

// A RefreshModeEnum is a refreshModeEnumType.
type RefreshModeEnum string

// RefreshModeEnums.
const (
	RefreshModeOnChange   RefreshModeEnum = "onChange"
	RefreshModeOnInterval RefreshModeEnum = "onInterval"
	RefreshModeOnExpire   RefreshModeEnum = "onExpire"
)

// A ViewRefreshModeEnum is a viewRefreshModeEnumType.
type ViewRefreshModeEnum string

// ViewRefreshModeEnums.
const (
	ViewRefreshModeNever     ViewRefreshModeEnum = "never"
	ViewRefreshModeOnRequest ViewRefreshModeEnum = "onRequest"
	ViewRefreshModeOnStop    ViewRefreshModeEnum = "onStop"
	ViewRefreshModeOnRegion  ViewRefreshModeEnum = "onRegion"
)

// A ShapeEnum is a shapeEnumType.
type ShapeEnum string

// ShapeEnums.
const (
	ShapeRectangle ShapeEnum = "rectangle"
	ShapeCylinder  ShapeEnum = "cylinder"
	ShapeSphere    ShapeEnum = "sphere"
)

// A StyleStateEnum is a styleStateEnumType.
type StyleStateEnum string

// StyleStateEnums.
const (
	StyleStateNormal    StyleStateEnum = "normal"
	StyleStateHighlight StyleStateEnum = "highlight"
)

// A UnitsEnum is a unitsEnumType.
type UnitsEnum string

// UnitsEnums.
const (
	UnitsFraction    UnitsEnum = "fraction"
	UnitsPixels      UnitsEnum = "pixels"
	UnitsInsetPixels UnitsEnum = "insetPixels"
)

// Address returns a new address element.
func Address(value string) *SimpleElement {
	return newSEString("address", value)
}

// Altitude returns a new altitude element.
func Altitude(value float64) *SimpleElement {
	return newSEFloat("altitude", value)
}

// AltitudeMode returns a new altitudeMode element.
func AltitudeMode(value AltitudeModeEnum) *SimpleElement {
	return newSEString("altitudeMode", string(value))
}

// Begin returns a new begin element.
func Begin(value time.Time) *SimpleElement {
	return newSETime("begin", value)
}

// BgColor returns a new bgColor element.
func BgColor(value color.Color) *SimpleElement {
	return newSEColor("bgColor", value)
}

// BottomFOV returns a new bottomFov element.
func BottomFOV(value float64) *SimpleElement {
	return newSEFloat("bottomFov", value)
}

// Color returns a new color element.
func Color(value color.Color) *SimpleElement {
	return newSEColor("color", value)
}

// ColorMode returns a new colorMode element.
func ColorMode(value ColorModeEnum) *SimpleElement {
	return newSEString("colorMode", string(value))
}

// Cookie returns a new cookie element.
func Cookie(value string) *SimpleElement {
	return newSEString("cookie", value)
}

// Description returns a new description element.
func Description(value string) *SimpleElement {
	return newSEString("description", value)
}

// DisplayName returns a new displayName element.
func DisplayName(value string) *SimpleElement {
	return newSEString("displayName", value)
}

// DisplayMode returns a new displayMode element.
func DisplayMode(value DisplayModeEnum) *SimpleElement {
	return newSEString("displayMode", string(value))
}

// DrawOrder returns a new drawOrder element.
func DrawOrder(value int) *SimpleElement {
	return newSEInt("drawOrder", value)
}

// East returns a new east element.
func East(value float64) *SimpleElement {
	return newSEFloat("east", value)
}

// End returns a new end element.
func End(value time.Time) *SimpleElement {
	return newSETime("end", value)
}

// Expires returns a new expires element.
func Expires(value time.Time) *SimpleElement {
	return newSETime("expires", value)
}

// Extrude returns a new extrude element.
func Extrude(value bool) *SimpleElement {
	return newSEBool("extrude", value)
}

// Fill returns a new fill element.
func Fill(value bool) *SimpleElement {
	return newSEBool("fill", value)
}

// FlyToView returns a new flyToView element.
func FlyToView(value bool) *SimpleElement {
	return newSEBool("flyToView", value)
}

// GridOrigin returns a new gridOrigin element.
func GridOrigin(value GridOriginEnum) *SimpleElement {
	return newSEString("gridOrigin", string(value))
}

// Heading returns a new heading element.
func Heading(value float64) *SimpleElement {
	return newSEFloat("heading", value)
}

// Href returns a new href element.
func Href(value string) *SimpleElement {
	return newSEString("href", value)
}

// HTTPQuery returns a new httpQuery element.
func HTTPQuery(value string) *SimpleElement {
	return newSEString("httpQuery", value)
}

// HotSpot returns a new hotSpot element.
func HotSpot(value Vec2) *SimpleElement {
	return newSEVec2("hotSpot", value)
}

// Key returns a new key element.
func Key(value StyleStateEnum) *SimpleElement {
	return newSEString("key", string(value))
}

// Latitude returns a new latitude element.
func Latitude(value float64) *SimpleElement {
	return newSEFloat("latitude", value)
}

// LeftFOV returns a new leftFov element.
func LeftFOV(value float64) *SimpleElement {
	return newSEFloat("leftFov", value)
}

// LinkDescription returns a new linkDescription element.
func LinkDescription(value string) *SimpleElement {
	return newSEString("linkDescription", value)
}

// LinkName returns a new linkName element.
func LinkName(value string) *SimpleElement {
	return newSEString("linkName", value)
}

// ListItemType returns a new listItemType element.
func ListItemType(value ListItemTypeEnum) *SimpleElement {
	return newSEString("listItemType", string(value))
}

// Longitude returns a new longitude element.
func Longitude(value float64) *SimpleElement {
	return newSEFloat("longitude", value)
}

// MaxSnippetLines returns a new maxSnippetLines element.
func MaxSnippetLines(value int) *SimpleElement {
	return newSEInt("maxSnippetLines", value)
}

// MaxSessionLength returns a new maxSessionLength element.
func MaxSessionLength(value float64) *SimpleElement {
	return newSEFloat("maxSessionLength", value)
}

// Message returns a new message element.
func Message(value string) *SimpleElement {
	return newSEString("message", value)
}

// MinAltitude returns a new minAltitude element.
func MinAltitude(value float64) *SimpleElement {
	return newSEFloat("minAltitude", value)
}

// MinFadeExtent returns a new minFadeExtent element.
func MinFadeExtent(value float64) *SimpleElement {
	return newSEFloat("minFadeExtent", value)
}

// MinLODPixels returns a new minLodPixels element.
func MinLODPixels(value float64) *SimpleElement {
	return newSEFloat("minLodPixels", value)
}

// MinRefreshPeriod returns a new minRefreshPeriod element.
func MinRefreshPeriod(value float64) *SimpleElement {
	return newSEFloat("minRefreshPeriod", value)
}

// MaxAltitude returns a new maxAltitude element.
func MaxAltitude(value float64) *SimpleElement {
	return newSEFloat("maxAltitude", value)
}

// MaxFadeExtent returns a new maxFadeExtent element.
func MaxFadeExtent(value float64) *SimpleElement {
	return newSEFloat("maxFadeExtent", value)
}

// MaxLODPixels returns a new maxLodPixels element.
func MaxLODPixels(value float64) *SimpleElement {
	return newSEFloat("maxLodPixels", value)
}

// MaxHeight returns a new maxHeight element.
func MaxHeight(value int) *SimpleElement {
	return newSEInt("maxHeight", value)
}

// MaxWidth returns a new maxWidth element.
func MaxWidth(value int) *SimpleElement {
	return newSEInt("maxWidth", value)
}

// Name returns a new name element.
func Name(value string) *SimpleElement {
	return newSEString("name", value)
}

// Near returns a new near element.
func Near(value float64) *SimpleElement {
	return newSEFloat("near", value)
}

// North returns a new north element.
func North(value float64) *SimpleElement {
	return newSEFloat("north", value)
}

// Open returns a new open element.
func Open(value bool) *SimpleElement {
	return newSEBool("open", value)
}

// Outline returns a new outline element.
func Outline(value bool) *SimpleElement {
	return newSEBool("outline", value)
}

// OverlayXY returns a new overlayXY element.
func OverlayXY(value Vec2) *SimpleElement {
	return newSEVec2("overlayXY", value)
}

// PhoneNumber returns a new phoneNumber element.
func PhoneNumber(value string) *SimpleElement {
	return newSEString("phoneNumber", value)
}

// Range returns a new range element.
func Range(value float64) *SimpleElement {
	return newSEFloat("range", value)
}

// RefreshMode returns a new refreshMode element.
func RefreshMode(value RefreshModeEnum) *SimpleElement {
	return newSEString("refreshMode", string(value))
}

// RefreshInterval returns a new refreshInterval element.
func RefreshInterval(value float64) *SimpleElement {
	return newSEFloat("refreshInterval", value)
}

// RefreshVisibility returns a new refreshVisibility element.
func RefreshVisibility(value bool) *SimpleElement {
	return newSEBool("refreshVisibility", value)
}

// RightFOV returns a new rightFov element.
func RightFOV(value float64) *SimpleElement {
	return newSEFloat("rightFov", value)
}

// Roll returns a new roll element.
func Roll(value float64) *SimpleElement {
	return newSEFloat("roll", value)
}

// Rotation returns a new rotation element.
func Rotation(value float64) *SimpleElement {
	return newSEFloat("rotation", value)
}

// RotationXY returns a new rotationXY element.
func RotationXY(value Vec2) *SimpleElement {
	return newSEVec2("rotationXY", value)
}

// Scale returns a new scale element.
func Scale(value float64) *SimpleElement {
	return newSEFloat("scale", value)
}

// ScreenXY returns a new screenXY element.
func ScreenXY(value Vec2) *SimpleElement {
	return newSEVec2("screenXY", value)
}

// Shape returns a new shape element.
func Shape(value ShapeEnum) *SimpleElement {
	return newSEString("shape", string(value))
}

// Size returns a new size element.
func Size(value Vec2) *SimpleElement {
	return newSEVec2("size", value)
}

// South returns a new south element.
func South(value float64) *SimpleElement {
	return newSEFloat("south", value)
}

// SourceHref returns a new sourceHref element.
func SourceHref(value string) *SimpleElement {
	return newSEString("sourceHref", value)
}

// Snippet returns a new snippet element.
func Snippet(value string) *SimpleElement {
	return newSEString("snippet", value)
}

// State returns a new state element.
func State(value ItemIconStateEnum) *SimpleElement {
	return newSEString("state", string(value))
}

// StyleURL returns a new styleUrl element.
func StyleURL(value string) *SimpleElement {
	return newSEString("styleUrl", value)
}

// TargetHref returns a new targetHref element.
func TargetHref(value string) *SimpleElement {
	return newSEString("targetHref", value)
}

// Tessellate returns a new tessellate element.
func Tessellate(value bool) *SimpleElement {
	return newSEBool("tessellate", value)
}

// Text returns a new text element.
func Text(value string) *SimpleElement {
	return newSEString("text", value)
}

// TextColor returns a new textColor element.
func TextColor(value color.Color) *SimpleElement {
	return newSEColor("textColor", value)
}

// TileSize returns a new tileSize element.
func TileSize(value int) *SimpleElement {
	return newSEInt("tileSize", value)
}

// Tilt returns a new tilt element.
func Tilt(value float64) *SimpleElement {
	return newSEFloat("tilt", value)
}

// TopFOV returns a new topFov element.
func TopFOV(value float64) *SimpleElement {
	return newSEFloat("topFov", value)
}

// Value returns a new value element.
func Value(value string) *SimpleElement {
	return newSEString("value", value)
}

// ViewBoundScale returns a new viewBoundScale element.
func ViewBoundScale(value float64) *SimpleElement {
	return newSEFloat("viewBoundScale", value)
}

// ViewFormat returns a new viewFormat element.
func ViewFormat(value string) *SimpleElement {
	return newSEString("viewFormat", value)
}

// ViewRefreshMode returns a new viewRefreshMode element.
func ViewRefreshMode(value ViewRefreshModeEnum) *SimpleElement {
	return newSEString("viewRefreshMode", string(value))
}

// ViewRefreshTime returns a new viewRefreshTime element.
func ViewRefreshTime(value float64) *SimpleElement {
	return newSEFloat("viewRefreshTime", value)
}

// Visibility returns a new visibility element.
func Visibility(value bool) *SimpleElement {
	return newSEBool("visibility", value)
}

// West returns a new west element.
func West(value float64) *SimpleElement {
	return newSEFloat("west", value)
}

// When returns a new when element.
func When(value time.Time) *SimpleElement {
	return newSETime("when", value)
}

// Width returns a new width element.
func Width(value float64) *SimpleElement {
	return newSEFloat("width", value)
}

// X returns a new x element.
func X(value float64) *SimpleElement {
	return newSEFloat("x", value)
}

// Y returns a new y element.
func Y(value float64) *SimpleElement {
	return newSEFloat("y", value)
}

// Z returns a new z element.
func Z(value float64) *SimpleElement {
	return newSEFloat("z", value)
}

// LookAt returns a new LookAt element.
func LookAt(children ...Element) *CompoundElement {
	return newCE("LookAt", children)
}

// Camera returns a new Camera element.
func Camera(children ...Element) *CompoundElement {
	return newCE("Camera", children)
}

// Metadata returns a new Metadata element.
func Metadata(children ...Element) *CompoundElement {
	return newCE("Metadata", children)
}

// ExtendedData returns a new ExtendedData element.
func ExtendedData(children ...Element) *CompoundElement {
	return newCE("ExtendedData", children)
}

// Data returns a new Data element.
func Data(children ...Element) *CompoundElement {
	return newCE("Data", children)
}

// NetworkLinkControl returns a new NetworkLinkControl element.
func NetworkLinkControl(children ...Element) *CompoundElement {
	return newCE("NetworkLinkControl", children)
}

// Document returns a new Document element.
func Document(children ...Element) *CompoundElement {
	return newCE("Document", children)
}

// Folder returns a new Folder element.
func Folder(children ...Element) *CompoundElement {
	return newCE("Folder", children)
}

// Placemark returns a new Placemark element.
func Placemark(children ...Element) *CompoundElement {
	return newCE("Placemark", children)
}

// NetworkLink returns a new NetworkLink element.
func NetworkLink(children ...Element) *CompoundElement {
	return newCE("NetworkLink", children)
}

// Region returns a new Region element.
func Region(children ...Element) *CompoundElement {
	return newCE("Region", children)
}

// LatLonAltBox returns a new LatLonAltBox element.
func LatLonAltBox(children ...Element) *CompoundElement {
	return newCE("LatLonAltBox", children)
}

// LOD returns a new Lod element.
func LOD(children ...Element) *CompoundElement {
	return newCE("Lod", children)
}

// Icon returns a new Icon element.
func Icon(children ...Element) *CompoundElement {
	return newCE("Icon", children)
}

// Link returns a new Link element.
func Link(children ...Element) *CompoundElement {
	return newCE("Link", children)
}

// URL returns a new Url element.
func URL(children ...Element) *CompoundElement {
	return newCE("Url", children)
}

// MultiGeometry returns a new MultiGeometry element.
func MultiGeometry(children ...Element) *CompoundElement {
	return newCE("MultiGeometry", children)
}

// Point returns a new Point element.
func Point(children ...Element) *CompoundElement {
	return newCE("Point", children)
}

// LineString returns a new LineString element.
func LineString(children ...Element) *CompoundElement {
	return newCE("LineString", children)
}

// LinearRing returns a new LinearRing element.
func LinearRing(children ...Element) *CompoundElement {
	return newCE("LinearRing", children)
}

// Polygon returns a new Polygon element.
func Polygon(children ...Element) *CompoundElement {
	return newCE("Polygon", children)
}

// OuterBoundaryIs returns a new outerBoundaryIs element.
func OuterBoundaryIs(value Element) *CompoundElement {
	return newSEElement("outerBoundaryIs", value)
}

// InnerBoundaryIs returns a new innerBoundaryIs element.
func InnerBoundaryIs(value Element) *CompoundElement {
	return newSEElement("innerBoundaryIs", value)
}

// Model returns a new Model element.
func Model(children ...Element) *CompoundElement {
	return newCE("Model", children)
}

// Location returns a new Location element.
func Location(children ...Element) *CompoundElement {
	return newCE("Location", children)
}

// Orientation returns a new Orientation element.
func Orientation(children ...Element) *CompoundElement {
	return newCE("Orientation", children)
}

// ResourceMap returns a new ResourceMap element.
func ResourceMap(children ...Element) *CompoundElement {
	return newCE("ResourceMap", children)
}

// Alias returns a new Alias element.
func Alias(children ...Element) *CompoundElement {
	return newCE("Alias", children)
}

// GroundOverlay returns a new GroundOverlay element.
func GroundOverlay(children ...Element) *CompoundElement {
	return newCE("GroundOverlay", children)
}

// LatLonBox returns a new LatLonBox element.
func LatLonBox(children ...Element) *CompoundElement {
	return newCE("LatLonBox", children)
}

// ScreenOverlay returns a new ScreenOverlay element.
func ScreenOverlay(children ...Element) *CompoundElement {
	return newCE("ScreenOverlay", children)
}

// PhotoOverlay returns a new PhotoOverlay element.
func PhotoOverlay(children ...Element) *CompoundElement {
	return newCE("PhotoOverlay", children)
}

// ViewVolume returns a new ViewVolume element.
func ViewVolume(children ...Element) *CompoundElement {
	return newCE("ViewVolume", children)
}

// ImagePyramid returns a new ImagePyramid element.
func ImagePyramid(children ...Element) *CompoundElement {
	return newCE("ImagePyramid", children)
}

// Style returns a new Style element.
func Style(children ...Element) *CompoundElement {
	return newCE("Style", children)
}

// StyleMap returns a new StyleMap element.
func StyleMap(children ...Element) *CompoundElement {
	return newCE("StyleMap", children)
}

// Pair returns a new Pair element.
func Pair(children ...Element) *CompoundElement {
	return newCE("Pair", children)
}

// IconStyle returns a new IconStyle element.
func IconStyle(children ...Element) *CompoundElement {
	return newCE("IconStyle", children)
}

// LabelStyle returns a new LabelStyle element.
func LabelStyle(children ...Element) *CompoundElement {
	return newCE("LabelStyle", children)
}

// LineStyle returns a new LineStyle element.
func LineStyle(children ...Element) *CompoundElement {
	return newCE("LineStyle", children)
}

// PolyStyle returns a new PolyStyle element.
func PolyStyle(children ...Element) *CompoundElement {
	return newCE("PolyStyle", children)
}

// BalloonStyle returns a new BalloonStyle element.
func BalloonStyle(children ...Element) *CompoundElement {
	return newCE("BalloonStyle", children)
}

// ListStyle returns a new ListStyle element.
func ListStyle(children ...Element) *CompoundElement {
	return newCE("ListStyle", children)
}

// ItemIcon returns a new ItemIcon element.
func ItemIcon(children ...Element) *CompoundElement {
	return newCE("ItemIcon", children)
}

// TimeStamp returns a new TimeStamp element.
func TimeStamp(children ...Element) *CompoundElement {
	return newCE("TimeStamp", children)
}

// TimeSpan returns a new TimeSpan element.
func TimeSpan(children ...Element) *CompoundElement {
	return newCE("TimeSpan", children)
}

// Update returns a new Update element.
func Update(children ...Element) *CompoundElement {
	return newCE("Update", children)
}

// Create returns a new Create element.
func Create(children ...Element) *CompoundElement {
	return newCE("Create", children)
}

// Delete returns a new Delete element.
func Delete(children ...Element) *CompoundElement {
	return newCE("Delete", children)
}

// Change returns a new Change element.
func Change(children ...Element) *CompoundElement {
	return newCE("Change", children)
}
