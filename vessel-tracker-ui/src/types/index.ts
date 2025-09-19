import { SVGProps } from "react";

export type IconSvgProps = SVGProps<SVGSVGElement> & {
  size?: number;
};

export type Position = {
  latitude: number;
  longitude: number;
};

export type Bounds = {
  maxLatitude: number;
  minLatitude: number;
  minLongitude: number;
  maxLongitude: number;
};
