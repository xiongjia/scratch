import React from 'react';
import Button from '@mui/material/Button';
import { Typography } from "@mui/material";

export interface DashBoardProps {
    name?: string
}

const Dashboard = ({name}: DashBoardProps) => {

    return (
        <div>
            <Typography>{name || 'default'}</Typography>
            <Button variant="contained">Dashboard</Button>
        </div>
    );
};


export default Dashboard;
