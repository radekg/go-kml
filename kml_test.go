package kml

import (
	"bytes"
	"encoding/xml"
	"image/color"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSimpleElements(t *testing.T) {
	for _, tc := range []struct {
		name     string
		element  Element
		expected string
	}{
		{
			name:     "Altitude",
			element:  Altitude(0),
			expected: `<altitude>0</altitude>`,
		},
		{
			name:     "AltitudeMode",
			element:  AltitudeMode(AltitudeModeAbsolute),
			expected: `<altitudeMode>absolute</altitudeMode>`,
		},
		{
			name:     "Begin",
			element:  Begin(time.Date(1876, 8, 1, 0, 0, 0, 0, time.UTC)),
			expected: `<begin>1876-08-01T00:00:00Z</begin>`,
		},
		{
			name:     "BgColor",
			element:  BgColor(color.Black),
			expected: `<bgColor>ff000000</bgColor>`,
		},
		{
			name:     "Color",
			element:  Color(color.White),
			expected: `<color>ffffffff</color>`,
		},
		{
			name:     "Coordinates",
			element:  Coordinates(Coordinate{Lon: 1.23, Lat: 4.56, Alt: 7.89}),
			expected: `<coordinates>1.23,4.56,7.89</coordinates>`,
		},
		{
			name:     "CoordinatesArray0",
			element:  CoordinatesArray([]float64{1.23, 4.56}),
			expected: `<coordinates>1.23,4.56</coordinates>`,
		},
		{
			name:     "CoordinatesArray1",
			element:  CoordinatesArray([]float64{1.23, 4.56, 7.89}),
			expected: `<coordinates>1.23,4.56,7.89</coordinates>`,
		},
		{
			name:     "CoordinatesArray2",
			element:  CoordinatesArray([][]float64{{1.23, 4.56}, {7.89, 0.12}}...),
			expected: `<coordinates>1.23,4.56 7.89,0.12</coordinates>`,
		},
		{
			element:  CoordinatesFlat([]float64{1.23, 4.56, 7.89, 0.12}, 0, 4, 2, 2),
			expected: `<coordinates>1.23,4.56 7.89,0.12</coordinates>`,
		},
		{
			name:     "CoordinatesFlat0",
			element:  CoordinatesFlat([]float64{1.23, 4.56, 0, 7.89, 0.12, 0}, 0, 6, 3, 3),
			expected: `<coordinates>1.23,4.56 7.89,0.12</coordinates>`,
		},
		{
			name:     "CoordinatesFlat1",
			element:  CoordinatesFlat([]float64{1.23, 4.56, 7.89, 0.12, 3.45, 6.78}, 0, 6, 3, 3),
			expected: `<coordinates>1.23,4.56,7.89 0.12,3.45,6.78</coordinates>`,
		},
		{
			name:     "Description",
			element:  Description("text"),
			expected: `<description>text</description>`,
		},
		{
			name:     "End",
			element:  End(time.Date(2015, 12, 31, 23, 59, 59, 0, time.UTC)),
			expected: `<end>2015-12-31T23:59:59Z</end>`,
		},
		{
			name:     "Extrude",
			element:  Extrude(false),
			expected: `<extrude>0</extrude>`,
		},
		{
			name:     "Folder",
			element:  Folder(),
			expected: `<Folder></Folder>`,
		},
		{
			name:     "GxCoord",
			element:  GxCoord(Coordinate{1.23, 4.56, 7.89}),
			expected: `<gx:coord>1.23 4.56 7.89</gx:coord>`,
		},
		{
			name:     "Heading",
			element:  Heading(0),
			expected: `<heading>0</heading>`,
		},
		{
			name:     "HotSpot",
			element:  HotSpot(Vec2{X: 0.5, Y: 0.5, XUnits: UnitsPixels, YUnits: UnitsPixels}),
			expected: `<hotSpot x="0.5" y="0.5" xunits="pixels" yunits="pixels"></hotSpot>`,
		},
		{
			name:     "Href",
			element:  Href("https://www.google.com/"),
			expected: `<href>https://www.google.com/</href>`,
		},
		{
			name:     "Latitude",
			element:  Latitude(0),
			expected: `<latitude>0</latitude>`,
		},
		{
			name:     "LinkSnippet",
			element:  LinkSnippet(2, "snippet"),
			expected: `<linkSnippet maxLines="2">snippet</linkSnippet>`,
		},
		{
			name:     "ListItemType",
			element:  ListItemType(ListItemTypeCheck),
			expected: `<listItemType>check</listItemType>`,
		},
		{
			name:     "OverlayXY",
			element:  OverlayXY(Vec2{X: 0, Y: 0, XUnits: UnitsFraction, YUnits: UnitsFraction}),
			expected: `<overlayXY x="0" y="0" xunits="fraction" yunits="fraction"></overlayXY>`,
		},
		{
			name:     "Style",
			element:  Style(),
			expected: `<Style></Style>`,
		},
		// FIXME More simple elements
	} {
		t.Run(tc.name, func(t *testing.T) {
			b := &bytes.Buffer{}
			e := xml.NewEncoder(b)
			require.NoError(t, e.Encode(tc.element))
			assert.Equal(t, tc.expected, b.String())
		})
	}
}

func TestCompoundElements(t *testing.T) {
	for _, tc := range []struct {
		name     string
		element  Element
		expected string
	}{
		{
			name: "easy_trail",
			element: Placemark(
				Name("Easy trail"),
				ExtendedData(
					SchemaData("#TrailHeadTypeId",
						SimpleData("TrailHeadName", "Pi in the sky"),
						SimpleData("TrailLength", "3.14159"),
						SimpleData("ElevationGain", "10"),
					),
				),
				Point(
					Coordinates(Coordinate{Lon: -122.000, Lat: 37.002}),
				),
			),
			expected: `<Placemark>` +
				`<name>Easy trail</name>` +
				`<ExtendedData>` +
				`<SchemaData schemaUrl="#TrailHeadTypeId">` +
				`<SimpleData name="TrailHeadName">Pi in the sky</SimpleData>` +
				`<SimpleData name="TrailLength">3.14159</SimpleData>` +
				`<SimpleData name="ElevationGain">10</SimpleData>` +
				`</SchemaData>` +
				`</ExtendedData>` +
				`<Point>` +
				`<coordinates>-122,37.002</coordinates>` +
				`</Point>` +
				`</Placemark>`,
		},
		{
			name: "simple_crosshairs",
			element: ScreenOverlay(
				Name("Simple crosshairs"),
				Description("This screen overlay uses fractional positioning to put the image in the exact center of the screen"),
				Icon(
					Href("http://myserver/myimage.jpg"),
				),
				OverlayXY(Vec2{X: 0.5, Y: 0.5, XUnits: UnitsFraction, YUnits: UnitsFraction}),
				ScreenXY(Vec2{X: 0.5, Y: 0.5, XUnits: UnitsFraction, YUnits: UnitsFraction}),
				Rotation(39.37878630116985),
				Size(Vec2{X: 0, Y: 0, XUnits: UnitsPixels, YUnits: UnitsPixels}),
			),
			expected: `<ScreenOverlay>` +
				`<name>Simple crosshairs</name>` +
				`<description>This screen overlay uses fractional positioning to put the image in the exact center of the screen</description>` +
				`<Icon>` +
				`<href>http://myserver/myimage.jpg</href>` +
				`</Icon>` +
				`<overlayXY x="0.5" y="0.5" xunits="fraction" yunits="fraction"></overlayXY>` +
				`<screenXY x="0.5" y="0.5" xunits="fraction" yunits="fraction"></screenXY>` +
				`<rotation>39.37878630116985</rotation>` +
				`<size x="0" y="0" xunits="pixels" yunits="pixels"></size>` +
				`</ScreenOverlay>`,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			b := &bytes.Buffer{}
			e := xml.NewEncoder(b)
			require.NoError(t, e.Encode(tc.element))
			assert.Equal(t, tc.expected, b.String())
		})
	}
}

func TestSharedStyles(t *testing.T) {
	style0 := SharedStyle("0")
	highlightPlacemarkStyle := SharedStyle(
		"highlightPlacemark",
		IconStyle(
			Icon(
				Href("http://maps.google.com/mapfiles/kml/paddle/red-stars.png"),
			),
		),
	)
	normalPlacemarkStyle := SharedStyle(
		"normalPlacemark",
		IconStyle(
			Icon(
				Href("http://maps.google.com/mapfiles/kml/paddle/wht-blank.png"),
			),
		),
	)
	exampleStyleMap := SharedStyleMap(
		"exampleStyleMap",
		Pair(
			Key(StyleStateNormal),
			StyleURL(normalPlacemarkStyle.URL()),
		),
		Pair(
			Key(StyleStateHighlight),
			StyleURL(highlightPlacemarkStyle.URL()),
		),
	)
	for _, tc := range []struct {
		name     string
		element  Element
		expected string
	}{
		{
			name: "folder",
			element: Folder(
				style0,
				Placemark(
					StyleURL(style0.URL()),
				),
			),
			expected: `<Folder>` +
				`<Style id="0">` +
				`</Style>` +
				`<Placemark>` +
				`<styleUrl>#0</styleUrl>` +
				`</Placemark>` +
				`</Folder>`,
		},
		{
			name: "highlighted_icon",
			element: KML(
				Document(
					Name("Highlighted Icon"),
					Description("Place your mouse over the icon to see it display the new icon"),
					highlightPlacemarkStyle,
					normalPlacemarkStyle,
					exampleStyleMap,
					Placemark(
						Name("Roll over this icon"),
						StyleURL(exampleStyleMap.URL()),
						Point(
							Coordinates(Coordinate{Lon: -122.0856545755255, Lat: 37.42243077405461}),
						),
					),
				),
			),
			expected: `<kml xmlns="http://www.opengis.net/kml/2.2">` +
				`<Document>` +
				`<name>Highlighted Icon</name>` +
				`<description>Place your mouse over the icon to see it display the new icon</description>` +
				`<Style id="highlightPlacemark">` +
				`<IconStyle>` +
				`<Icon>` +
				`<href>http://maps.google.com/mapfiles/kml/paddle/red-stars.png</href>` +
				`</Icon>` +
				`</IconStyle>` +
				`</Style>` +
				`<Style id="normalPlacemark">` +
				`<IconStyle>` +
				`<Icon>` +
				`<href>http://maps.google.com/mapfiles/kml/paddle/wht-blank.png</href>` +
				`</Icon>` +
				`</IconStyle>` +
				`</Style>` +
				`<StyleMap id="exampleStyleMap">` +
				`<Pair>` +
				`<key>normal</key>` +
				`<styleUrl>#normalPlacemark</styleUrl>` +
				`</Pair>` +
				`<Pair>` +
				`<key>highlight</key>` +
				`<styleUrl>#highlightPlacemark</styleUrl>` +
				`</Pair>` +
				`</StyleMap>` +
				`<Placemark>` +
				`<name>Roll over this icon</name>` +
				`<styleUrl>#exampleStyleMap</styleUrl>` +
				`<Point>` +
				`<coordinates>-122.0856545755255,37.42243077405461</coordinates>` +
				`</Point>` +
				`</Placemark>` +
				`</Document>` +
				`</kml>`,
		},
		{
			name: "trail_head_type",
			element: KML(
				Document(
					Schema("TrailHeadTypeId", "TrailHeadType",
						SimpleField("TrailHeadName", "string",
							DisplayName("<b>Trail Head Name</b>"),
						),
						SimpleField("TrailLength", "double",
							DisplayName("<i>The length in miles</i>"),
						),
						SimpleField("ElevationGain", "int",
							DisplayName("<i>change in altitude</i>"),
						),
					),
				),
			),
			expected: `<kml xmlns="http://www.opengis.net/kml/2.2">` +
				`<Document>` +
				`<Schema id="TrailHeadTypeId" name="TrailHeadType">` +
				`<SimpleField name="TrailHeadName" type="string">` +
				`<displayName>&lt;b&gt;Trail Head Name&lt;/b&gt;</displayName>` +
				`</SimpleField>` +
				`<SimpleField name="TrailLength" type="double">` +
				`<displayName>&lt;i&gt;The length in miles&lt;/i&gt;</displayName>` +
				`</SimpleField>` +
				`<SimpleField name="ElevationGain" type="int">` +
				`<displayName>&lt;i&gt;change in altitude&lt;/i&gt;</displayName>` +
				`</SimpleField>` +
				`</Schema>` +
				`</Document>` +
				`</kml>`,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			b := &bytes.Buffer{}
			e := xml.NewEncoder(b)
			require.NoError(t, e.Encode(tc.element))
			assert.Equal(t, tc.expected, b.String())
		})
	}
}

func TestWrite(t *testing.T) {
	for _, tc := range []struct {
		name     string
		element  Element
		expected string
	}{
		{
			name:    "placemark",
			element: KML(Placemark()),
			expected: `<?xml version="1.0" encoding="UTF-8"?>` + "\n" +
				`<kml xmlns="http://www.opengis.net/kml/2.2">` +
				`<Placemark>` +
				`</Placemark>` +
				`</kml>`,
		},
		{
			name: "simple_placemark",
			element: KML(
				Placemark(
					Name("Simple placemark"),
					Description("Attached to the ground. Intelligently places itself at the height of the underlying terrain."),
					Point(
						Coordinates(Coordinate{Lon: -122.0822035425683, Lat: 37.42228990140251}),
					),
				),
			),
			expected: `<?xml version="1.0" encoding="UTF-8"?>` + "\n" +
				`<kml xmlns="http://www.opengis.net/kml/2.2">` +
				`<Placemark>` +
				`<name>Simple placemark</name>` +
				`<description>Attached to the ground. Intelligently places itself at the height of the underlying terrain.</description>` +
				`<Point>` +
				`<coordinates>-122.0822035425683,37.42228990140251</coordinates>` +
				`</Point>` +
				`</Placemark>` +
				`</kml>`,
		},
		{
			name: "entity_references_example",
			element: KML(
				Document(
					Placemark(
						Name("Entity references example"),
						Description(
							`<h1>Entity references are hard to type!</h1>`+
								`<p><font color="red">Text is <i>more readable</i> and `+
								`<b>easier to write</b> when you can avoid using entity `+
								`references.</font></p>`,
						),
						Point(
							Coordinates(Coordinate{Lon: 102.594411, Lat: 14.998518}),
						),
					),
				),
			),
			expected: `<?xml version="1.0" encoding="UTF-8"?>` + "\n" +
				`<kml xmlns="http://www.opengis.net/kml/2.2">` +
				`<Document>` +
				`<Placemark>` +
				`<name>Entity references example</name>` +
				`<description>` +
				`&lt;h1&gt;Entity references are hard to type!&lt;/h1&gt;` +
				`&lt;p&gt;&lt;font color=&#34;red&#34;&gt;Text is ` +
				`&lt;i&gt;more readable&lt;/i&gt; ` +
				`and &lt;b&gt;easier to write&lt;/b&gt; ` +
				`when you can avoid using entity references.&lt;/font&gt;&lt;/p&gt;` +
				`</description>` +
				`<Point>` +
				`<coordinates>102.594411,14.998518</coordinates>` +
				`</Point>` +
				`</Placemark>` +
				`</Document>` +
				`</kml>`,
		},
		{
			name: "ground_overlays",
			element: KML(
				Folder(
					Name("Ground Overlays"),
					Description("Examples of ground overlays"),
					GroundOverlay(
						Name("Large-scale overlay on terrain"),
						Description("Overlay shows Mount Etna erupting on July 13th, 2001."),
						Icon(
							Href("http://developers.google.com/kml/documentation/images/etna.jpg"),
						),
						LatLonBox(
							North(37.91904192681665),
							South(37.46543388598137),
							East(15.35832653742206),
							West(14.60128369746704),
							Rotation(-0.1556640799496235),
						),
					),
				),
			),
			expected: `<?xml version="1.0" encoding="UTF-8"?>` + "\n" +
				`<kml xmlns="http://www.opengis.net/kml/2.2">` +
				`<Folder>` +
				`<name>Ground Overlays</name>` +
				`<description>Examples of ground overlays</description>` +
				`<GroundOverlay>` +
				`<name>Large-scale overlay on terrain</name>` +
				`<description>Overlay shows Mount Etna erupting on July 13th, 2001.</description>` +
				`<Icon>` +
				`<href>http://developers.google.com/kml/documentation/images/etna.jpg</href>` +
				`</Icon>` +
				`<LatLonBox>` +
				`<north>37.91904192681665</north>` +
				`<south>37.46543388598137</south>` +
				`<east>15.35832653742206</east>` +
				`<west>14.60128369746704</west>` +
				`<rotation>-0.1556640799496235</rotation>` +
				`</LatLonBox>` +
				`</GroundOverlay>` +
				`</Folder>` +
				`</kml>`,
		},
		{
			name: "the_pentagon",
			element: KML(
				Placemark(
					Name("The Pentagon"),
					Polygon(
						Extrude(true),
						AltitudeMode(AltitudeModeRelativeToGround),
						OuterBoundaryIs(
							LinearRing(
								Coordinates([]Coordinate{
									{-77.05788457660967, 38.87253259892824, 100},
									{-77.05465973756702, 38.87291016281703, 100},
									{-77.05315536854791, 38.87053267794386, 100},
									{-77.05552622493516, 38.868757801256, 100},
									{-77.05844056290393, 38.86996206506943, 100},
									{-77.05788457660967, 38.87253259892824, 100},
								}...),
							),
						),
						InnerBoundaryIs(
							LinearRing(
								Coordinates([]Coordinate{
									{-77.05668055019126, 38.87154239798456, 100},
									{-77.05542625960818, 38.87167890344077, 100},
									{-77.05485125901024, 38.87076535397792, 100},
									{-77.05577677433152, 38.87008686581446, 100},
									{-77.05691162017543, 38.87054446963351, 100},
									{-77.05668055019126, 38.87154239798456, 100},
								}...),
							),
						),
					),
				),
			),
			expected: `<?xml version="1.0" encoding="UTF-8"?>` + "\n" +
				`<kml xmlns="http://www.opengis.net/kml/2.2">` +
				`<Placemark>` +
				`<name>The Pentagon</name>` +
				`<Polygon>` +
				`<extrude>1</extrude>` +
				`<altitudeMode>relativeToGround</altitudeMode>` +
				`<outerBoundaryIs>` +
				`<LinearRing>` +
				`<coordinates>` +
				`-77.05788457660967,38.87253259892824,100 ` +
				`-77.05465973756702,38.87291016281703,100 ` +
				`-77.0531553685479,38.87053267794386,100 ` +
				`-77.05552622493516,38.868757801256,100 ` +
				`-77.05844056290393,38.86996206506943,100 ` +
				`-77.05788457660967,38.87253259892824,100` +
				`</coordinates>` +
				`</LinearRing>` +
				`</outerBoundaryIs>` +
				`<innerBoundaryIs>` +
				`<LinearRing>` +
				`<coordinates>` +
				`-77.05668055019126,38.87154239798456,100 ` +
				`-77.05542625960818,38.87167890344077,100 ` +
				`-77.05485125901023,38.87076535397792,100 ` +
				`-77.05577677433152,38.87008686581446,100 ` +
				`-77.05691162017543,38.87054446963351,100 ` +
				`-77.05668055019126,38.87154239798456,100` +
				`</coordinates>` +
				`</LinearRing>` +
				`</innerBoundaryIs>` +
				`</Polygon>` +
				`</Placemark>` +
				`</kml>`,
		},
		{
			name:    "gx_placemark",
			element: GxKML(Placemark()),
			expected: `<?xml version="1.0" encoding="UTF-8"?>` + "\n" +
				`<kml xmlns="http://www.opengis.net/kml/2.2" xmlns:gx="http://www.google.com/kml/ext/2.2">` +
				`<Placemark>` +
				`</Placemark>` +
				`</kml>`,
		},
		{
			name: "gx_track",
			element: GxKML(
				Folder(
					Placemark(
						GxTrack(
							When(time.Date(2010, 5, 28, 2, 2, 9, 0, time.UTC)),
							When(time.Date(2010, 5, 28, 2, 2, 35, 0, time.UTC)),
							When(time.Date(2010, 5, 28, 2, 2, 44, 0, time.UTC)),
							When(time.Date(2010, 5, 28, 2, 2, 53, 0, time.UTC)),
							When(time.Date(2010, 5, 28, 2, 2, 54, 0, time.UTC)),
							When(time.Date(2010, 5, 28, 2, 2, 55, 0, time.UTC)),
							When(time.Date(2010, 5, 28, 2, 2, 56, 0, time.UTC)),
							GxCoord(Coordinate{-122.207881, 37.371915, 156.000000}),
							GxCoord(Coordinate{-122.205712, 37.373288, 152.000000}),
							GxCoord(Coordinate{-122.204678, 37.373939, 147.000000}),
							GxCoord(Coordinate{-122.203572, 37.374630, 142.199997}),
							GxCoord(Coordinate{-122.203451, 37.374706, 141.800003}),
							GxCoord(Coordinate{-122.203329, 37.374780, 141.199997}),
							GxCoord(Coordinate{-122.203207, 37.374857, 140.199997}),
						),
					),
				),
			),
			expected: `<?xml version="1.0" encoding="UTF-8"?>` + "\n" +
				`<kml xmlns="http://www.opengis.net/kml/2.2" xmlns:gx="http://www.google.com/kml/ext/2.2">` +
				`<Folder>` +
				`<Placemark>` +
				`<gx:Track>` +
				`<when>2010-05-28T02:02:09Z</when>` +
				`<when>2010-05-28T02:02:35Z</when>` +
				`<when>2010-05-28T02:02:44Z</when>` +
				`<when>2010-05-28T02:02:53Z</when>` +
				`<when>2010-05-28T02:02:54Z</when>` +
				`<when>2010-05-28T02:02:55Z</when>` +
				`<when>2010-05-28T02:02:56Z</when>` +
				`<gx:coord>-122.207881 37.371915 156</gx:coord>` +
				`<gx:coord>-122.205712 37.373288 152</gx:coord>` +
				`<gx:coord>-122.204678 37.373939 147</gx:coord>` +
				`<gx:coord>-122.203572 37.37463 142.199997</gx:coord>` +
				`<gx:coord>-122.203451 37.374706 141.800003</gx:coord>` +
				`<gx:coord>-122.203329 37.37478 141.199997</gx:coord>` +
				`<gx:coord>-122.203207 37.374857 140.199997</gx:coord>` +
				`</gx:Track>` +
				`</Placemark>` +
				`</Folder>` +
				`</kml>`,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			b := &bytes.Buffer{}
			require.NoError(t, tc.element.Write(b))
			assert.Equal(t, tc.expected, b.String())
		})
	}
}
