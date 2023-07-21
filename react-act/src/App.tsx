import React, { useState } from 'react';
import MainSidebar from './components/MainSidebar';
import AppBar from './components/AppBar';
import Dashboard from './components/Dashboard';
import { Routes, Route } from 'react-router-dom';
import { Box } from '@mui/material';
import { PageContext } from './PageContext';
import useScrollTrigger from '@mui/material/useScrollTrigger';

const val = {
  name: 'test context',
};

interface Props {
  /**
   * Injected by the documentation to work in an iframe.
   * You won't need it on your project.
   */
  window?: () => Window;
  children: React.ReactElement;
}

function ElevationScroll(props: Props) {
  const { children, window } = props;
  // Note that you normally won't need to set the window ref as useScrollTrigger
  // will default to window.
  // This is only being set here because the demo is in an iframe.
  const trigger = useScrollTrigger({
    disableHysteresis: true,
    threshold: 0,
    target: window ? window() : undefined,
  });

  return React.cloneElement(children, {
    elevation: trigger ? 4 : 0,
  });
}

function App() {
  const [name, setName] = useState('default context');

  return (
    <PageContext.Provider value={{ name, setName }}>
      <div className="app">
        <ElevationScroll>
          <AppBar />
        </ElevationScroll>
        <main className="content" style={{ display: 'flex' }}>
          <MainSidebar />
          <Box flexGrow={1}>
            <Routes>
              <Route path="/" element={<Dashboard name="home" />} />
              <Route path="/data1" element={<Dashboard name="data1" />} />
              <Route path="/data2" element={<Dashboard name="data2" />} />
            </Routes>
          </Box>
        </main>
      </div>
    </PageContext.Provider>
  );
}

export default App;
