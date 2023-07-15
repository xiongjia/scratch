import React, { useContext } from 'react';
import Button from '@mui/material/Button';
import { Typography } from "@mui/material";
import { usePage } from '../PageContext';

export interface DashBoardProps {
    name?: string
}

const Dashboard = ({name}: DashBoardProps) => {
    const ctx = usePage();

    return (
        <div>
            <Typography>{name || 'default'} : {ctx.name}</Typography>
            <Button variant="contained" onClick={() => {
                console.log("click");
                ctx.setName("abc");
            }}>Dashboard</Button>
        </div>
    );
};


export default Dashboard;
