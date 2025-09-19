import { Route, Routes } from "react-router-dom";

import MapPage from "@/pages/mapPage.tsx";

function App() {
  return (
    <Routes>
      <Route element={<MapPage />} path="/" />
    </Routes>
  );
}

export default App;
