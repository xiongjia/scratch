import React from 'react';
import { useState } from "react";
import Sidebar from './components/MainSidebar';
import Button from '@mui/material/Button';

function App() {
  const [isSidebar, setIsSidebar] = useState(true);

  return (
    <div className='app'>
      <main className="content" style={{ display: "flex" }}>
        <Sidebar />
      </main>
    </div>
  );
}

export default App;
