import { Marker, Popup, TileLayer } from "react-leaflet";
import { useMap } from "react-leaflet/hooks";
import { useEffect, useState } from "react";

import { Bounds, Ship } from "@/types";
import  ShipComponent, { ShipComponentProps } from "./geojson/shipComponent";
import { fetchShips } from "@/api/ship";

export interface TileLayerProps {
  initialPosition: number[];
}

export default function TileLayerComponent(props: TileLayerProps) {
  const [bounds, setBounds] = useState<Bounds>();
  const [ships, setShips] = useState<Ship[]>([]);
  const map = useMap();

  useEffect(() => {
    async function updateBounds() {
      const currentBounds = map.getBounds();

      const bound: Bounds = {
        maxLatitude: currentBounds.getNorth(),
        minLatitude: currentBounds.getSouth(),
        maxLongitude: currentBounds.getEast(),
        minLongitude: currentBounds.getWest(),
      };
      setBounds(bound);
      setShips(await fetchShips(bound))
    }
    updateBounds();
    map.on("moveend", updateBounds);
    

    return () => {
      map.off("moveend", updateBounds);
    };
  }, [map]);

  const testShipProps: ShipComponentProps = {
    mmsi: 219016555,
    name: "LOLLAND",
    position: { latitude: 54.90533333333334, longitude: 10.900533333333334 },
    heading: 107,
  };

  return (
    <>
      <TileLayer
        attribution='&copy; <a href="http://osm.org/copyright">OpenStreetMap</a> contributors'
        url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
      />

      {ships.map((ship) => (
        <ShipComponent 
          mmsi={ship.mmsi}
          name={ship.name}
          position={ship.position}
          heading={ship.heading}
        ></ShipComponent>
      ))}

      <Marker position={props.initialPosition}>
        <Popup>
          <div className="absolute top-2 left-2 bg-white p-2 rounded shadow">
            {bounds ? (
              <>
                <p>Max Lat: {bounds.maxLatitude.toFixed(5)}</p>
                <p>Min Lat: {bounds.minLatitude.toFixed(5)}</p>
                <p>Max Lng: {bounds.maxLongitude.toFixed(5)}</p>
                <p>Min Lng: {bounds.minLongitude.toFixed(5)}</p>
              </>
            ) : (
              <p>Loading bounds...</p>
            )}
          </div>
        </Popup>
        `
      </Marker>
    </>
  );
}
