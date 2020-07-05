import { useState, useEffect } from 'react';
import react, * as React from "react";
import {Grid, StatsCard} from 'tabler-react';

function generateResult(input) {
  if (input === undefined) {
    return 0
  } else {
    return input
  }
}


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

export function ListEmployeeActiveEmployee() {
  const [stats, handleStats] = useState([]);

  const FetchData = async () => {
    const data = await fetch(process.env.REACT_APP_GATEWAY_URL + "/management/search/status");
    const stats = await data.json();
    handleStats(stats)  
  }

  useEffect(() => {
    FetchData()
  }, [])
  const empData = stats

  return (
    <Grid.Col sm={3}>
    <StatsCard
      layout={1}
      movement={0}
      total={generateResult(empData["Current Employee"])}
      label="Active Employees"
    />
    </Grid.Col>
  )
}

export function ListEmployeeInActiveEmployee() {
  const [stats, handleStats] = useState([]);

  const FetchData = async () => {
    const data = await fetch(process.env.REACT_APP_GATEWAY_URL + "/management/search/status");
    const stats = await data.json();
    handleStats(stats)  
  }

  useEffect(() => {
    FetchData()
  }, [])
  const empData = stats

  return (
    <Grid.Col sm={3}>
    <StatsCard
      layout={1}
      movement={0}
      total={generateResult(empData["Ex-Employee"])}
      label="Ex-Employees"
    />
    </Grid.Col>
  )
}
