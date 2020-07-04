import react, * as React from "react";
import { Page, Grid, Table, Button } from "tabler-react";
import SiteWrapper from "./SiteWrapper.react";

function ListAttendance() {
    return (
        <SiteWrapper>
        <Page.Card
              title="Attendance List"
          ></Page.Card>
          <Grid.Col md={6} lg={8} className="align-self-center">
          <Table>
            <Table.Header>
                <Table.ColHeader>Employee ID</Table.ColHeader>
                <Table.ColHeader>State</Table.ColHeader>
                <Table.ColHeader>Date</Table.ColHeader>
            </Table.Header>
            <Table.Body>
            <Table.Row>
                <Table.Col>1</Table.Col>
                <Table.Col>Present</Table.Col>
                <Table.Col>2020-07-07</Table.Col>
            </Table.Row>
            <Table.Row>
                <Table.Col>2</Table.Col>
                <Table.Col>Absent</Table.Col>
                <Table.Col>2020-07-07</Table.Col>
            </Table.Row>
            </Table.Body>
        </Table>
          </Grid.Col>
      </SiteWrapper>
    );
}

export default ListAttendance
