import react, * as React from "react";
import { Page, Card, Grid, Button, Form, Dropdown } from "tabler-react";
import SiteWrapper from "./SiteWrapper.react";
import { Formik } from 'formik';
import './EmployeForm.css';


const FormElements = () => (
  <SiteWrapper>
         <Page.Card
            title="Employee Registration"
            // RootComponent={Form}
        ></Page.Card>
        <Grid.Col md={6} lg={6} className="align-self-center">
  <div>
    <Formik
    initialValues={{ email: "" }}
    onSubmit={async values => {
      await new Promise(resolve => setTimeout(resolve, 500));
      alert(JSON.stringify(values, null, 2));
    }}
    >
    {props => {
        const {
          values,
          touched,
          errors,
          dirty,
          isSubmitting,
          handleChange,
          handleBlur,
          handleSubmit,
          handleReset
        } = props;
        return (
          <form onSubmit={handleSubmit}>
            <label htmlFor="id" style={{ display: "block" }}>
              ID
            </label>
            <input
              id="id"
              placeholder="Employee ID"
              type="text"
              onChange={handleChange}
              onBlur={handleBlur}
            />

            <label htmlFor="name" style={{ display: "block" }}>
              Name
            </label>
            <input
              id="name"
              placeholder="Employee Name"
              type="text"
              onChange={handleChange}
              onBlur={handleBlur}
            />

            <label htmlFor="job_role" style={{ display: "block" }}>
              Job Role
            </label>
            <select
              name="job_role"
              value={values.job_role}
              onChange={handleChange}
              onBlur={handleBlur}
              style={{ display: 'block' }}
            >
              <option value="red" label="red" />
              <option value="blue" label="blue" />
              <option value="green" label="green" />
            </select>

            <button
              type="button"
              className="outline"
              onClick={handleReset}
              disabled={!dirty || isSubmitting}
            >
              Reset
            </button>
            <button type="submit" disabled={isSubmitting}>
              Submit
            </button>
          </form>
        );
      }}
  </Formik>
  </div>
  </Grid.Col>
  </SiteWrapper>
);

export default FormElements
