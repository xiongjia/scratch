import React from 'react';
import { useState, createContext } from "react";
import MainSidebar from './components/MainSidebar';
import AppBar from './components/AppBar';
import Dashboard from './components/Dashboard';
import { Routes, Route } from "react-router-dom";
import { Box } from "@mui/material";

function App() {
  return (
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
  );
}

export default App;
