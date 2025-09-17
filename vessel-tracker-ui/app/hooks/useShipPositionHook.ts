import { useCallback, useEffect, useState } from 'react';
import {Ship} from "@/app/types/ship";

interface ShipPositionReturn {
    ships: Ship[];
    error: string | null;
    refetch: () => Promise<void>;
}



export const useShipPositionHook = (interval: number = 30000): ShipPositionReturn => {



}, []