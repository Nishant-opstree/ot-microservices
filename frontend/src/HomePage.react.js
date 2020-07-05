// @flow

import * as React from "react";

import {
  Page,
  Avatar,
  Icon,
  Grid,
  Card,
  Text,
  Table,
  Alert,
  Progress,
  colors,
  Dropdown,
  Button,
  StampCard,
  StatsCard,
  ProgressCard,
  Badge,
} from "tabler-react";

import C3Chart from "react-c3js";

import SiteWrapper from "./SiteWrapper.react";
import { ListAllEmployees } from './EmployeeData';

function Home() {
  return (
    <SiteWrapper>
      <Page.Content title="Dashboard">
        <Grid.Row cards={true}>
          {/* <Grid.Col sm={3}>
            <StatsCard layout={1} movement={0} total="150" label="Total Employees" />
          </Grid.Col> */}
          <ListAllEmployees/>
          <Grid.Col sm={3}>
            <StatsCard
              layout={1}
              movement={0}
              total="130"
              label="Active Employees"
            />
          </Grid.Col>
          <Grid.Col sm={3}>
            <StatsCard
              layout={1}
              movement={0}
              total="20"
              label="Ex-Employees"
            />
          </Grid.Col>
          <Grid.Col sm={3}>
            <StatsCard
              layout={1}
              movement={0}
              total="5"
              label="Office Locations"
            />
          </Grid.Col>
          <Grid.Col>
            <Grid.Row cards="true">
              <Grid.Col sm={4}>
                <Card>
                  <Card.Header>
                    <Card.Title>Job Role Distribution</Card.Title>
                  </Card.Header>
                  <Card.Body>
                    <C3Chart
                      style={{ height: "12rem" }}
                      data={{
                        columns: [
                          // each columns data
                          ["DevOps", 100],
                          ["Developer", 30],
                        ],
                        type: "donut", // default type of chart
                        colors: {
                          data1: colors["green"],
                          data2: colors["green-light"],
                        },
                        names: {
                          // name of each serie
                          data1: "Maximum",
                          data2: "Minimum",
                        },
                      }}
                      legend={{
                        show: false, //hide legend
                      }}
                      padding={{
                        bottom: 0,
                        top: 0,
                      }}
                    />
                  </Card.Body>
                </Card>
              </Grid.Col>
              <Grid.Col sm={4}>
                <Card>
                  <Card.Header>
                    <Card.Title>Locations Distribution</Card.Title>
                  </Card.Header>
                  <Card.Body>
                    <C3Chart
                      style={{ height: "12rem" }}
                      data={{
                        columns: [
                          // each columns data
                          ["Delhi", 63],
                          ["Bangalore", 44],
                          ["Hyederabad", 12],
                          ["Newyork", 14],
                        ],
                        type: "donut", // default type of chart
                        colors: {
                          data1: colors["blue-darker"],
                          data2: colors["blue"],
                          data3: colors["blue-light"],
                          data4: colors["blue-lighter"],
                        },
                        names: {
                          // name of each serie
                          data1: "A",
                          data2: "B",
                          data3: "C",
                          data4: "D",
                        },
                      }}
                      legend={{
                        show: false, //hide legend
                      }}
                      padding={{
                        bottom: 0,
                        top: 0,
                      }}
                    />
                  </Card.Body>
                </Card>
              </Grid.Col>
              <Grid.Col sm={4}>
                <Card>
                  <Card.Header>
                    <Card.Title>Employees Distribution</Card.Title>
                  </Card.Header>
                  <Card.Body>
                    <C3Chart
                      style={{ height: "12rem" }}
                      data={{
                        columns: [
                          // each columns data
                          ["Current Employees", 130],
                          ["Ex-Employees", 20],
                        ],
                        type: "donut", // default type of chart
                        colors: {
                          data1: colors["blue-darker"],
                          data2: colors["blue"],
                          data3: colors["blue-light"],
                          data4: colors["blue-lighter"],
                        },
                        names: {
                          // name of each serie
                          data1: "A",
                          data2: "B",
                          data3: "C",
                          data4: "D",
                        },
                      }}
                      legend={{
                        show: false, //hide legend
                      }}
                      padding={{
                        bottom: 0,
                        top: 0,
                      }}
                    />
                  </Card.Body>
                </Card>
              </Grid.Col>
            </Grid.Row>
          </Grid.Col>
        </Grid.Row>
      </Page.Content>
    </SiteWrapper>
  );
}

export default Home;