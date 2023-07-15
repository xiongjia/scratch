import React, { useState } from 'react';
import MainSidebar from './components/MainSidebar';
import AppBar from './components/AppBar';
import Dashboard from './components/Dashboard';
import { Routes, Route } from "react-router-dom";
import { Box } from "@mui/material";
import { PageContext } from "./PageContext";

const val = {
  name: 'test context' 
};

function App() {
  const [name, setName] = useState("default context");

  return (
      <PageContext.Provider value={{name, setName}}>
      <div className='app'>
        <AppBar />
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
    </PageContext.Provider>
  );
}

export default App;
