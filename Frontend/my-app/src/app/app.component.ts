import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css'],
})
export class AppComponent implements OnInit {
  public headsetInfo: any;
  public constructor(private http: HttpClient) { }
  public ngOnInit(): void {
    const url: string = 'http://localhost:4201/headsets';
    this.http.get(url).subscribe((response) => {
      this.headsetInfo = response;
    });
  }
}
