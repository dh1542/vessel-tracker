import { Position } from "@/types";
import { Circle, Popup } from "react-leaflet";

export interface ShipProps {
    mmsi: number;
    name: string;
    position: Position;
    heading: number;
}


export default function Ship(props: ShipProps){
    return (
        <>
            <Circle center={[props.position.latitude, props.position.longitude]} radius={50} pathOptions={{color: 'blue'}}>
                <Popup>
                    <div>
                        <h3>{props.name}</h3>
                        <p>MMSI: {props.mmsi}</p>
                        <p>Lat: {props.position.latitude.toFixed(5)}</p>
                        <p>Lng: {props.position.longitude.toFixed(5)}</p>
                    </div>
                </Popup>
            </Circle>
        </>
    )
}