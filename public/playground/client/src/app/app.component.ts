import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { ButtonModule } from 'primeng/button';
import { FormatJsonPipe } from './pipes/format-json.pipe';
import { AccordionModule } from 'primeng/accordion';
import { InputNumberModule } from 'primeng/inputnumber';
import { InputTextModule } from 'primeng/inputtext';
import { CheckboxModule } from 'primeng/checkbox';
import { map } from 'rxjs';
import { FormsModule } from '@angular/forms';

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
    CheckboxModule,
    FormsModule, // Add FormsModule here
  ],
  templateUrl: './app.component.html',
  styleUrl: './app.component.scss',
})
export class AppComponent implements OnInit {
  endpoints: Endpoint[] = [];

  constructor(private httpClient: HttpClient) {}
  ngOnInit(): void {
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
    if (endpoint.method === 'POST') {
      this.httpClient
        .post('http://localhost:8080' + endpoint.path, endpoint.body)
        .subscribe((data) => {
          console.log(data);
          endpoint.response = data;
        });
    } else if (endpoint.method === 'PUT') {
      this.httpClient
        .put('http://localhost:8080' + endpoint.path, endpoint.body)
        .subscribe((data) => {
          console.log(data);
          endpoint.response = data;
        });
    } else if (endpoint.method === 'DELETE') {
      this.httpClient
        .delete('http://localhost:8080' + endpoint.path)
        .subscribe((data) => {
          console.log(data);
          endpoint.response = data;
        });
    } else {
      this.httpClient
        .request(endpoint.method, 'http://localhost:8080' + endpoint.path)
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
}

type Endpoint = {
  name: string;
  method: string;
  path: string;
  response?: any;
  body: any;
  fields?: { key: string; type: string }[];
};
