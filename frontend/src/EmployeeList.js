import react, * as React from "react";
import { Page, Grid, Table, Button } from "tabler-react";
import SiteWrapper from "./SiteWrapper.react";

function ListEmployee() {
    return (
        <SiteWrapper>
        <Page.Card
              title="Employee List"
          ></Page.Card>
          <Grid.Col md={6} lg={8} className="align-self-center">
          <Table>
            <Table.Header>
                <Table.ColHeader>Employee ID</Table.ColHeader>
                <Table.ColHeader>Name</Table.ColHeader>
                <Table.ColHeader>Email</Table.ColHeader>
                <Table.ColHeader>Phone Number</Table.ColHeader>
                <Table.ColHeader>Job Role</Table.ColHeader>
                <Table.ColHeader>Job Location</Table.ColHeader>
            </Table.Header>
            <Table.Body>
            <Table.Row>
                <Table.Col>1</Table.Col>
                <Table.Col>Abhishek</Table.Col>
                <Table.Col>abhishek.dubey@opstree.com</Table.Col>
                <Table.Col>9560375769</Table.Col>
                <Table.Col>DevOps</Table.Col>
                <Table.Col>Delhi</Table.Col>
            </Table.Row>
            <Table.Row>
                <Table.Col>2</Table.Col>
                <Table.Col>Vishant</Table.Col>
                <Table.Col>vishant.sharma@opstree.com</Table.Col>
                <Table.Col>9999999999</Table.Col>
                <Table.Col>DevOps</Table.Col>
                <Table.Col>Hyderabad</Table.Col>
            </Table.Row>
            </Table.Body>
        </Table>
          </Grid.Col>
      </SiteWrapper>
    );
}

export default ListEmployee
