import { HttpClient, HttpHeaders } from '@angular/common/http';
import { AfterViewInit, Component, OnInit } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { ButtonModule } from 'primeng/button';
import { FormatJsonPipe } from './pipes/format-json.pipe';
import { AccordionModule } from 'primeng/accordion';
import { InputNumberModule } from 'primeng/inputnumber';
import { InputTextModule } from 'primeng/inputtext';
import { CheckboxModule } from 'primeng/checkbox';
import { catchError, map, throwError } from 'rxjs';
import { FormsModule } from '@angular/forms';
import { MenubarModule } from 'primeng/menubar';
import { ChipModule } from 'primeng/chip';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [
    RouterOutlet,
    ButtonModule,
    FormatJsonPipe,
    AccordionModule,
    InputNumberModule,
    InputTextModule,
    MenubarModule,
    ChipModule,
    CheckboxModule,
    FormsModule, // Add FormsModule here
  ],
  templateUrl: './app.component.html',
  styleUrl: './app.component.scss',
})
export class AppComponent implements OnInit,AfterViewInit {
  endpoints: Endpoint[] = [];
  email = "tylor@email.com";
  password = "Test123456";
  token = "";

  constructor(private httpClient: HttpClient) {}
  ngAfterViewInit(): void {
    this.login();
  }
  ngOnInit(): void {
    this.login();
    this.httpClient
      .get('http://localhost:8080/endpoints')
      .pipe(
        map((endpoints) => {
          return (endpoints as Endpoint[]).map((e: Endpoint) => {
            if (e.body) {
              const fields = [];
              for (const key in e.body) {
                fields.push({
                  key,
                  type: this.getType(
                    typeof e.body[key] as 'string' | 'number' | 'boolean'
                  ),
                });
              }
              e.fields = fields;
            }

            return e as Endpoint;
          });
        })
      )
      .subscribe((data) => {
        console.log(data);
        this.endpoints = data as Endpoint[];
      });
  }
  runRoute(endpoint: Endpoint): void {
    const httpOptions = {
      headers: new HttpHeaders({
        'Authorization': `Bearer ${this.token}`
      })
    };
    if (endpoint.method === 'POST') {
      this.httpClient
        .post('http://localhost:8080' + endpoint.path, {
          body: endpoint.body,
          ...httpOptions
        })
        .pipe(
          catchError((error) => {
            console.log(error);
            endpoint.error = error.error;
            return error;
          })
        )
        .subscribe((data) => {
          console.log(data);
          endpoint.response = data;
        });
    } else if (endpoint.method === 'PUT') {
      this.httpClient
        .put('http://localhost:8080' + endpoint.path, {
          body: endpoint.body,
          ...httpOptions
        })
        .pipe(
          catchError((error) => {
            console.log(error);
            endpoint.error = error.error.error;
            return error;
          })
        )
        .subscribe((data) => {
          console.log(data);
          endpoint.response = data;
        });
    } else if (endpoint.method === 'DELETE') {
      this.httpClient
        .delete('http://localhost:8080' + endpoint.path, {
          ...httpOptions
        })
        .pipe(
          catchError((error) => {
            console.log(error);
            endpoint.error = error.error.error;
            return error;
          })
        )
        .subscribe((data) => {
          console.log(data);
          endpoint.response = data;
        });
    } else {
      console.log(">>>", httpOptions.headers.getAll('Authorization'));
      this.httpClient
        .request(endpoint.method, 'http://localhost:8080' + endpoint.path, {
          body: endpoint.body,
          headers:  httpOptions.headers,
        })
        .pipe(
          catchError((error) => {
            console.log(error);
            endpoint.error = error.error.error;
            return throwError(() => new Error(error));
          })
        )
        .subscribe((data) => {
          console.log(data);
          endpoint.response = data;
        });
    }
  }

  private getType(type: 'string' | 'number' | 'boolean') {
    switch (type) {
      case 'string':
        return 'text';
      case 'number':
        return 'number';
      case 'boolean':
        return 'checkbox';
      default:
        return 'text';
    }
  }

  login() {
    this.httpClient
      .post('http://localhost:8080/login', {
        email: this.email,
        password: this.password,
      })
      .pipe(
        catchError((error) => {
          console.log(error);
          return error;
        })
      )
      .subscribe((data: any) => {
        this.token = data.token;
        console.log(data);
      });
  }
}



type Endpoint = {
  name: string;
  method: string;
  path: string;
  error: any;
  response?: any;
  body: any;
  fields?: { key: string; type: string }[];
};
