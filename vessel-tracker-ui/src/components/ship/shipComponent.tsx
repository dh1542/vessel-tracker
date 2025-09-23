import { Position } from "@/types";
import { Circle, Polyline, Popup } from "react-leaflet";
import ShipPopup from "./shipPopUp";

export interface ShipComponentProps {
  mmsi: number;
  name: string;
  position: Position;
  heading: number;
  destination?: string;
}

// React component to render the arrow shape
function ArrowIcon({ heading }: { heading: number }) {
  return (
    <div
      style={{
        display: "flex",
        alignItems: "center",
        justifyContent: "center",
        transform: `rotate(${heading}deg)`,
      }}
    >
      {/* Shaft */}
      <div
        style={{
          width: "2px",
          height: "20px",
          background: "blue",
        }}
      />
      {/* Arrowhead */}
      <div
        style={{
          width: 0,
          height: 0,
          borderLeft: "6px solid transparent",
          borderRight: "6px solid transparent",
          borderTop: "10px solid blue",
          marginLeft: "-4px",
        }}
      />
    </div>
  );
}

export default function ShipComponent(props: ShipComponentProps) {
  const { latitude, longitude } = props.position;
  const distance = 0.001;
  const rad = (props.heading * Math.PI) / 180;
  const endLat = latitude + distance * Math.cos(rad);
  const endLng = longitude + distance * Math.sin(rad);

  return (
    <>
      <Circle
        center={[latitude, longitude]}
        radius={50}
        pathOptions={{ color: "blue" }}
        onClick={() => {
          console.log(`Ship ${props.name} clicked`);
        }}
      >
        <ShipPopup {...props} />
      </Circle>
      <Polyline
        positions={[
          [latitude, longitude],
          [endLat, endLng],
        ]}
        pathOptions={{ color: "blue" }}
      />
    </>
  );
}
