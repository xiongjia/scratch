import React from 'react';
import { useState, createContext } from "react";
import MainSidebar from './components/MainSidebar';
import Dashboard from './components/Dashboard';
import { BrowserRouter, Routes, Route } from "react-router-dom";
import { Box } from "@mui/material";

function App() {
  const [isSidebar, setIsSidebar] = useState(true);
  const context = createContext({});


  return (
    <BrowserRouter>
      <div className='app'>
      <main className="content" style={{ display: "flex" }}>
        <MainSidebar/>
        <Box flexGrow={1}>
          <Routes>
            <Route path="/" element={<Dashboard name="home"/>} />
            <Route path="/data1" element={<Dashboard name="data1"/>} />
            <Route path="/data2" element={<Dashboard name="data2" />} />
          </Routes>
        </Box>
      </main>
    </div>
    </BrowserRouter>
  );
}

export default App;
