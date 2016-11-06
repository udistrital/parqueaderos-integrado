// format used to parse WFS GetFeature responses
var geojsonFormat = new ol.format.GeoJSON()

var islaSource = new ol.source.Vector({
  loader: function(extent, resolution, projection) {
    var url = '/geoserver/parqueaderos/ows?service=WFS&' +
      'version=1.0.0&request=GetFeature&typename=parqueaderos:isla&' +
      'outputFormat=application%2Fjson' +
      '&srsname=EPSG:3857&bbox=' + extent.join(',') + ',EPSG:3857'
      // use jsonp: false to prevent jQuery from adding the "callback"
      // parameter to the URL
    $.ajax({
      url: url,
      dataType: 'json',
      jsonp: false
    }).done(function(response) {
      islaSource.addFeatures(geojsonFormat.readFeatures(response))
    })
  },
  strategy: ol.loadingstrategy.tile(ol.tilegrid.createXYZ({
    maxZoom: 19
  }))
})

var grupoislaSource = new ol.source.Vector({
  loader: function(extent, resolution, projection) {
    var url = '/geoserver/parqueaderos/ows?service=WFS&' +
      'version=1.0.0&request=GetFeature&typename=parqueaderos:grupo_isla&' +
      'outputFormat=application%2Fjson&' +
      'srsname=EPSG:3857&bbox=' + extent.join(',') + ',EPSG:3857'
      // use jsonp: false to prevent jQuery from adding the "callback"
      // parameter to the URL
    $.ajax({
      url: url,
      dataType: 'json',
      jsonp: false
    }).done(function(response) {
      islaSource.addFeatures(geojsonFormat.readFeatures(response))
    })
  },
  strategy: ol.loadingstrategy.tile(ol.tilegrid.createXYZ({
    maxZoom: 19
  }))
})

var islaLayer = new ol.layer.Vector({
  source: islaSource,
  style: new ol.style.Style({
    stroke: new ol.style.Stroke({
      color: 'rgba(0, 0, 255, 1.0)',
      width: 2
    })
  })
})

var grupoislaLayer = new ol.layer.Vector({
  source: grupoislaSource,
  style: new ol.style.Style({
    stroke: new ol.style.Stroke({
      color: 'rgba(255, 0, 0, 1.0)',
      width: 2
    })
  })
})

var wmsSource = new ol.source.TileWMS({
  url: '/geoserver/wms',
  params: {
    'LAYERS': 'parqueaderos:sotano1_50ppp_georeferenced',
    'FORMAT': 'image/png',
    'TILED': true
  },
  serverType: 'geoserver',
  crossOrigin: 'anonymous'
})

var wmsLayer = new ol.layer.Tile({
  source: wmsSource
})

var osmLayer = new ol.layer.Tile({
  source: new ol.source.OSM()
})

//document.body.onclick = function (e) {console.log(map.getEventCoordinate(e))}
var map = new ol.Map({
  layers: [osmLayer, wmsLayer, grupoislaLayer, islaLayer],
  target: document.getElementById('map'),
  view: new ol.View({
    center: [-8244952.014276695, 515728.84084898204],
    maxZoom: 26,
    zoom: 18
  })
})

map.on('moveend', checknewzoom)

function checknewzoom(evt) {
  var newZoomLevel = map.getView().getZoom()
  if (newZoomLevel > 20) {
    wmsLayer.setVisible(true)
  } else {
    wmsLayer.setVisible(false)
  }
}

var sidebar = $('#sidebar').sidebar()

function changeVisibility(elem) {
  elem.estado = (elem.estado === true) ? false : true
  var estado = elem.estado
  var opt = $(elem).attr('data-layer')
  changeVisibilityLayer(opt, estado)
  var elem = $(elem).children()[0]
  changeVisibilityElement(elem, estado)
}

function changeVisibilityElement(elem, estado) {
  console.log(elem, estado)
  if (estado) {
    $(elem).text('visibility_on')
  } else {
    $(elem).text('visibility_off')
  }
}

function changeVisibilityLayer(opt, estado) {
  console.log(opt, estado)
  switch (opt) {
    case 'sotano1_raster':
      wmsLayer.setVisible(estado)
      break;
    case 'sotano1_vector':
      grupoislaLayer.setVisible(estado)
      break;
    case 'sotano1_islas_vector':
      islaLayer.setVisible(estado)
      break;
    default:

  }
}
