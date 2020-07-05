import { useState, useEffect } from 'react';
import react, * as React from "react";
import {Grid, StatsCard} from 'tabler-react';

export function ListAllEmployees() {
    const [stats, handleStats] = useState([]);

    const FetchData = async () => {
      const data = await fetch(process.env.REACT_APP_GATEWAY_URL + "/management/search/all");
      const stats = await data.json();
      handleStats(stats)  
    }
  
    useEffect(() => {
      FetchData()
    }, [])
    const empData = stats.length
    
    return (
      <Grid.Col sm={3}>
        <StatsCard 
            layout={1} 
            movement={0} 
            total={empData} 
            label="Total Employees" 
        />
      </Grid.Col>
    )
}
