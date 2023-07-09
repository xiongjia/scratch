import React from 'react';
import { useState } from "react";
import { Sidebar, Menu, SubMenu, MenuItem } from "react-pro-sidebar";
import { Box, IconButton, Typography } from "@mui/material";
import HomeOutlinedIcon from "@mui/icons-material/HomeOutlined";
import MenuOutlinedIcon from "@mui/icons-material/MenuOutlined";
import { Link } from "react-router-dom";

const MainSidebar = () => {
  const [isCollapsed, setIsCollapsed] = useState(false);

  return (
    <Box>
      <Sidebar collapsed={isCollapsed}>
        <Menu>
          <MenuItem icon={isCollapsed ? <MenuOutlinedIcon /> : undefined}
            onClick={() => setIsCollapsed(!isCollapsed)}
            style={{ margin: "10px 0 20px 0" }}>
            {!isCollapsed && (
              <Box display="flex"
                justifyContent="space-between"
                alignItems="center"
                ml="15px">
                <Typography>DashBoard</Typography>
                <IconButton onClick={() => setIsCollapsed(!isCollapsed)}>
                  <MenuOutlinedIcon />
                </IconButton>
              </Box>
            )}
          </MenuItem>
          <SubMenu label="Home" icon={<HomeOutlinedIcon />} defaultOpen >
            <MenuItem> Data1 <Link to="/data1" /> </MenuItem>
            <MenuItem> Data2 <Link to="/data2" /></MenuItem>
          </SubMenu>
        </Menu>
      </Sidebar>
    </Box>
  );
};

export default MainSidebar;
