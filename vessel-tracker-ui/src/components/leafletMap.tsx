import {
  MapContainer,
  Marker,
  Popup,
  TileLayer,
  TileLayerProps,
} from "react-leaflet";
import {  useState } from "react";

import { Bounds } from "@/types";
import TileLayerComponent from "@/components/tileLayer.tsx";

export default function LeafletMap() {

  const initialPosition: number[] = [53.528034753637016, 9.92970943450928];
  



  return (
    <MapContainer
      center={initialPosition}
      zoom={15}
      scrollWheelZoom={true}
      style={{ minHeight: "100vh", minWidth: "100vw" }}
    >
      <TileLayerComponent initialPosition={initialPosition}></TileLayerComponent>
    </MapContainer>
  );
}
