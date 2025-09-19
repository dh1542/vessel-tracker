import LeafletMap from "@/components/leafletMap.tsx";
import NavigationBar from "@/components/navigationBar";

export default function MapPage() {
  console.log("MapPage");
  return (
    <>
      <NavigationBar/>
      <LeafletMap/>
    </>
    
  );
}