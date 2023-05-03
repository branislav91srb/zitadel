import { Injectable } from '@angular/core';
import { Request, UnaryInterceptor, UnaryResponse } from 'grpc-web';
import { ExhaustedService } from '../exhausted.service';

@Injectable({ providedIn: 'root' })
/**
 * Show authenticated requests exhausted dialog if the cookie is present after the request
 */
export class ExhaustedInterceptor<TReq = unknown, TResp = unknown> implements UnaryInterceptor<TReq, TResp> {
  constructor(private exhaustedService: ExhaustedService) {}

  public async intercept(request: Request<TReq, TResp>, invoker: any): Promise<UnaryResponse<TReq, TResp>> {
    return invoker(request)
      .then((response: any) => {
        this.exhaustedService.checkCookie()
        return response;
      })
  }
}
