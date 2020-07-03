import * as React from "react";
import { Page, Card, Grid, Button, Form, Dropdown } from "tabler-react";
import SiteWrapper from "./SiteWrapper.react";
import { useFormik } from 'formik'
// import { Form, Input, SubmitBtn, Select, Datepicker } from 'react-formik-ui';


function FormElements() {
    const formik = useFormik({
      initialValues: {
        email: '',
      },
      onSubmit: values => {
        alert(JSON.stringify(values, null, 2));
      },
    });
    return (
        <SiteWrapper>
        <Page.Card
            title="Employee Registration"
            // RootComponent={Form}
        ></Page.Card>
        <Grid.Col md={6} lg={6} className="align-self-center">
            {/* <form onSubmit={formik.handleSubmit}>
            <label htmlFor="email">Email Address </label>
            <input
              id="email"
              name="email"
              type="email"
              onChange={formik.handleChange}
              value={formik.values.email}
            />
            <button type="submit">Submit</button>
          </form> */}
          <Form onSubmit={formik.handleSubmit}>
          <Form.Input name="email"  label="Email"/>
          <button type="submit">Submit</button>
          </Form>
        </Grid.Col>
      </SiteWrapper>
    );
}

// const FormElements = () => {
//       return (
//         <SiteWrapper>
//         <Page.Card
//             title="Employee Registration"
//             // RootComponent={Form}
//         ></Page.Card>
//         <Grid.Col md={6} lg={6} className="align-self-center">
//           <Formik
//             initialValues={{
//               name: '',
//             }}
//             onSubmit={data =>
//               // fetch('http://172.17.0.3:8080/create', {
//               //   method: 'POST',
//               //   body: JSON.stringify(data),
//               //   headers: {
//               //       'Content-Type': 'application/json'
//               // }})
//               (alert(JSON.stringify(data)))
//             }
//           >
//             <Form>
//               <Input
//                 name='name'
//                 label='Name'
//                 placeholder='Abhishek'
//                 required
//               />
//                <SubmitBtn />
//             </Form>
//           </Formik>
//         </Grid.Col>
//         </SiteWrapper>
//       )
// }

export default FormElements
// function FormElements() {
//     return (
//         <SiteWrapper>
//             <Page.Card
//                 title="Employee Registration"
//                 RootComponent={Form}
//             ></Page.Card>
//             {/* onSubmit={event =>
//                 // fetch('http://172.17.0.3:8080/create', {
//                 // method: 'POST',
//                 // body: JSON.stringify(data),
//                 // headers: {
//                 //     'Content-Type': 'application/json'
//                 // }})
//                 (alert(JSON.stringify(event)))
//             } */}
//             <Grid.Col md={6} lg={6} className="align-self-center">
//             <Form onSubmit={(event) => console.log(event.target.name + 'clicked')}>
//             <Form.Input name="name"  label="Name"/>
//         {/* <Form.FieldSet>
//             <Form.Group
//                 isRequired
//                 label="Name"
//             >
//                 <Form.Input name="name" />
//             </Form.Group>
//             <Form.Group
//                 isRequired
//                 label="Id"
//             >
//                 <Form.Input name="id" />
//             </Form.Group>
//             <Form.Group
//                 isRequired
//                 label="Email"
//             >
//                 <Form.Input name="email_id" />
//             </Form.Group>
//             <Form.Group
//                 className="mb-0"
//                 label="Phone number"
//             >
//                 <Form.Input name="phone_number" />
//             </Form.Group>
//             <Form.Group label="Location" name="location">
//             <Form.Select>
//                 <option>
//                 Delhi
//                 </option>
//                 <option>
//                 Bangalore
//                 </option>
//                 <option>
//                 Newyork
//                 </option>
//                 <option>
//                 Hyderabad
//                 </option>
//             </Form.Select>
//             </Form.Group>
//             <Form.Group label="Status" name="status">
//             <Form.Select>
//                 <option>
//                 Current Employee
//                 </option>
//                 <option>
//                 Ex-Employee
//                 </option>
//             </Form.Select>
//             </Form.Group>
//             <Form.Group label="Job Role" name="job_role">
//             <Form.Select>
//                 <option>
//                 Developer
//                 </option>
//                 <option>
//                 DevOps
//                 </option>
//             </Form.Select>
//             </Form.Group>
//             <Form.Group
//                 className="mb-0"
//                 label="Address"
//             >
//                 <Form.Input name="address" />
//             </Form.Group>
//             <Form.Group
//                 className="mb-0"
//                 label="Annual Package"
//             >
//                 <Form.Input name="annual_package" type="number" />
//             </Form.Group>
//             <Button type="submit" color="primary" className="ml-auto" md="1">
//                 Submit
//             </Button>
//         </Form.FieldSet> */}
//         {/* <SubmitBtn /> */}
//             <Button type="submit" color="primary" className="ml-auto" md="1">
//                 Submit
//             </Button>
//           </Form>
//         </Grid.Col>
//         </SiteWrapper>
//     );
// }

// export default FormElements