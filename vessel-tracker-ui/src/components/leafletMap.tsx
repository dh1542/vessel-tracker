import {
  MapContainer,
  Marker,
  Popup,
  TileLayer,
  TileLayerProps,
} from "react-leaflet";
import { useMap } from "react-leaflet/hooks";
import { useEffect, useState } from "react";

import { Bounds } from "@/types";
import TileLayerComponent from "@/components/tileLayer.tsx";

export default function LeafletMap() {
  const [bounds, setBounds] = useState<Bounds>();

  const initialPosition: number[] = [53.505, 10];


  // const currentBounds = map.getCenter();
  //
  // console.log("current bounds", currentBounds);
  //
  // useEffect(() => {
  //   function updateBounds() {
  //
  //
  //   }
  // })


  return (
    <MapContainer
      center={initialPosition}
      zoom={13}
      scrollWheelZoom={true}
      style={{ minHeight: "100vh", minWidth: "100vw" }}
    >
      <TileLayerComponent initialPosition={initialPosition}></TileLayerComponent>
    </MapContainer>
  );
}
