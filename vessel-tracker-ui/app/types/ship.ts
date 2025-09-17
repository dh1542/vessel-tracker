export interface Ship {
    Mmsi: bigint;
    ShipName: string;
    Latitude: number;
    Longitude: number;
    Cog: number;
    Sog: number;
    TrueHeading: number;
    NavigationalStatus: number;
    PositionAccuracy: boolean;
    CommunicationState: number;
    RateOfTurn: number;
    SpecialManoeuvreIndicator: number;
    RepeatIndicator: number;
    MessageId: number;
    Valid: boolean;
    TimeUtc: number;
}