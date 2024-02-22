var mymap = L.map('map').setView([47.2, -1.55], 5);

var markerLondon = L.marker([51.505, -0.09]).addTo(mymap);
markerLondon.bindPopup("<b>Londres: 30 février gazo</b>").openPopup();

var markerNantes = L.marker([47.2, -1.55]).addTo(mymap);
markerNantes.bindPopup("<b>Nantes</b>").openPopup();

var markerParis = L.marker([48.8566, 2.3522]).addTo(mymap);
markerParis.bindPopup("<b>Paris</b>").openPopup();

var markerBerlin = L.marker([52.5200, 13.4050]).addTo(mymap);
markerBerlin.bindPopup("<b>Berlin: ACDC le 31 Février.</b><br><img src='acdc logo.jpeg' width ='200'>").openPopup();

L.tileLayer('https://%7Bs%7D.tile.openstreetmap.org/%7Bz%7D/%7Bx%7D/%7By%7D.png', {
    attribution: '© OpenStreetMap contributors'
}).addTo(mymap);