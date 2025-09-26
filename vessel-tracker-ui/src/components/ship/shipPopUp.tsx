import { Popup } from "react-leaflet";
import { ShipComponentProps } from "./shipComponent";
import { Card, CardHeader, CardBody, Image } from "@heroui/react";
import { getFirstGoogleImage } from "@/api/imageFetcher";
import ShipPopUpContent from "./shipPopUpContent";

export default function ShipPopup(props: ShipComponentProps): JSX.Element {
  return (
    <Popup >
      {/* <div className="w-200 h-100p-2">
        <h3 className="text-lg font-semibold">{props.name}</h3>
        <p className="text-sm">MMSI: {props.mmsi}</p>
        <p className="text-sm">Lat: {props.position.latitude.toFixed(5)}</p>
        <p className="text-sm">Lng: {props.position.longitude.toFixed(5)}</p>
        <p className="text-sm">Destination: {props.destination}</p>
      </div> */}

      <ShipPopUpContent {...props} />
    </Popup>
  );
}
