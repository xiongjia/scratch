import React from 'react';
import { useState } from "react";
import { Sidebar, Menu, SubMenu, MenuItem } from "react-pro-sidebar";
import { Box } from "@mui/material";

const MainSidebar = () => {
  const [isCollapsed, setIsCollapsed] = useState(false);
  const [selected, setSelected] = useState("dashboard");

  return (
    <Box
      sx={{
        "& .pro-icon-wrapper": {
          backgroundColor: "transparent !important",
        },
        "& .pro-inner-item": {
          padding: "5px 35px 5px 20px !important",
        },
        "& .pro-inner-item:hover": {
          color: "#868dfb !important",
        },
        "& .pro-menu-item.active": {
          color: "#6870fa !important",
        },
      }}
    >
      <Sidebar>
        <Menu
          menuItemStyles={{
            button: {
              [`&.active`]: { backgroundColor: 'red', color: '#b6c8d9', },
            },
          }}
        >
          <SubMenu label="Charts">
            <MenuItem> Pie charts </MenuItem>
            <MenuItem> Line charts </MenuItem>
          </SubMenu>
          <MenuItem> Documentation </MenuItem>
          <MenuItem> Calendar </MenuItem>
        </Menu>
      </Sidebar>
    </Box>
  );
};

export default MainSidebar;
