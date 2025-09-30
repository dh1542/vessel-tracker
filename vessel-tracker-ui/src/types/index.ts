import { SVGProps } from "react";

export type IconSvgProps = SVGProps<SVGSVGElement> & {
  size?: number;
};

export type Position = {
  latitude: number;
  longitude: number;
};

export type SQLNonNullString = {
  string: string;
  valid: boolean;
};

export type Bounds = {
  maxLatitude: number;
  minLatitude: number;
  minLongitude: number;
  maxLongitude: number;
};

export type Ship = {
  mmsi: number;
  name: string;
  position: Position;
  heading: number;
  destination: string;
  imageUrl: string;
};
