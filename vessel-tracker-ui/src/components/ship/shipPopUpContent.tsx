import {
  Card,
  Tab,
  Table,
  TableBody,
  TableCell,
  TableColumn,
  TableHeader,
  TableRow,
} from "@heroui/react";
import { ShipComponentProps } from "./shipComponent";

export default function ShipPopUpContent(
  props: ShipComponentProps
): JSX.Element {
  return (
    <div className="w-bg-white shadow-2xs">
      <Card fullWidth={true} shadow="lg">
        <Table>
          <TableHeader>
            <TableColumn>NAME</TableColumn>
            <TableColumn>VALUE</TableColumn>
          </TableHeader>
          <TableBody>
            <TableRow key="1">
              <TableCell>Ship Name</TableCell>
              <TableCell>{props.name}</TableCell>
            </TableRow>
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
      </Card>
    </div>
  );
}
