import {
  Card,
  Tab,
  Table,
  TableBody,
  TableCell,
  TableColumn,
  TableHeader,
  TableRow,
  Image,
  Divider,
} from "@heroui/react";
import { ShipComponentProps } from "./shipComponent";

export default function ShipPopUpContent(
  props: ShipComponentProps
): JSX.Element {
  return (
    <div className="w-full bg-white">
      <h2 className="text-blue-500 text-xl font-semibold">{props.name}</h2>
      <Image
        alt="Ship image"
        src={props.imageUrl}
        className="pt-3"
        width={400}
      />
      <Divider className="my-4" />
      <Table className="w-full">
        <TableHeader>
          <TableColumn>NAME</TableColumn>
          <TableColumn>VALUE</TableColumn>
        </TableHeader>
        <TableBody>
          <TableRow key="2">
            <TableCell>Destination</TableCell>
            <TableCell>{props.destination}</TableCell>
          </TableRow>
          <TableRow key="3">
            <TableCell>MMSI</TableCell>
            <TableCell>{props.mmsi}</TableCell>
          </TableRow>
          <TableRow key="4">
            <TableCell>Latitude</TableCell>
            <TableCell>{props.position.latitude}</TableCell>
          </TableRow>
          <TableRow key="5">
            <TableCell>Longitude</TableCell>
            <TableCell>{props.position.longitude}</TableCell>
          </TableRow>
        </TableBody>
      </Table>
    </div>
  );
}
