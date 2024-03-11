import { Pipe, PipeTransform } from '@angular/core';

@Pipe({
  name: 'formatJson',
  standalone: true
})
export class FormatJsonPipe implements PipeTransform {

  transform(value: unknown, ...args: unknown[]): unknown {
    return JSON.stringify(value, null, 2);
  }
}
