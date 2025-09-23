import { ShipComponentProps } from "./shipComponent";
import { Card, CardHeader, CardBody, Image } from "@heroui/react";

export default function ShipPopup(props: ShipComponentProps): JSX.Element {
  return (
    <Card className="py-4">
      <div>
        <h3>{props.name}</h3>
        <p>MMSI: {props.mmsi}</p>
        <p>Lat: {props.position.latitude.toFixed(5)}</p>
        <p>Lng: {props.position.longitude.toFixed(5)}</p>
        <p>Destination: {props.destination}</p>
      </div>
    </Card>
  );
}
