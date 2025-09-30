import { Popup } from "react-leaflet";
import { ShipComponentProps } from "./shipComponent";
import { Card, CardHeader, CardBody, Image } from "@heroui/react";
import { getFirstGoogleImage } from "@/api/imageFetcher";
import ShipPopUpContent from "./shipPopUpContent";

export default function ShipPopup(props: ShipComponentProps): JSX.Element {
  return (
    <Popup>
      <ShipPopUpContent {...props} />
    </Popup>
  );
}
