// Créer une carte OpenStreetMap
var map = L.map('map').setView([47.20567535891242, -1.5393572608259531], 15); // Coordonnées initiales et niveau de zoo

// Ajouter une couche de tuiles OpenStreetMap
L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
    attribution: '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
}).addTo(map);