/* tslint:disable */
/* eslint-disable */
/**
 * Devmarks API
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 1.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import * as runtime from '../runtime';
import {
    LoginRequest,
    LoginRequestFromJSON,
    LoginRequestToJSON,
    RegisterRequest,
    RegisterRequestFromJSON,
    RegisterRequestToJSON,
    RegisterResponse,
    RegisterResponseFromJSON,
    RegisterResponseToJSON,
    Token,
    TokenFromJSON,
    TokenToJSON,
    User,
    UserFromJSON,
    UserToJSON,
} from '../models';

export interface GetUserRequest {
    embed?: string;
}

export interface LoginOperationRequest {
    loginRequest: LoginRequest;
}

export interface RegisterOperationRequest {
    registerRequest: RegisterRequest;
}

/**
 * 
 */
export class UserApi extends runtime.BaseAPI {

    /**
     * User Endpoint, returns the user corresponding to the supplied bearer token
     */
    async getUserRaw(requestParameters: GetUserRequest): Promise<runtime.ApiResponse<User>> {
        const queryParameters: any = {};

        if (requestParameters.embed !== undefined) {
            queryParameters['embed'] = requestParameters.embed;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.accessToken) {
            const token = this.configuration.accessToken;
            const tokenString = await token("bearerAuth", []);

            if (tokenString) {
                headerParameters["Authorization"] = `Bearer ${tokenString}`;
            }
        }
        const response = await this.request({
            path: `/me`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => UserFromJSON(jsonValue));
    }

    /**
     * User Endpoint, returns the user corresponding to the supplied bearer token
     */
    async getUser(requestParameters: GetUserRequest): Promise<User> {
        const response = await this.getUserRaw(requestParameters);
        return await response.value();
    }

    /**
     * authenticates a user based on the request body and returns an authentication bearer token
     */
    async loginRaw(requestParameters: LoginOperationRequest): Promise<runtime.ApiResponse<Token>> {
        if (requestParameters.loginRequest === null || requestParameters.loginRequest === undefined) {
            throw new runtime.RequiredError('loginRequest','Required parameter requestParameters.loginRequest was null or undefined when calling login.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/auth/token`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: LoginRequestToJSON(requestParameters.loginRequest),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => TokenFromJSON(jsonValue));
    }

    /**
     * authenticates a user based on the request body and returns an authentication bearer token
     */
    async login(requestParameters: LoginOperationRequest): Promise<Token> {
        const response = await this.loginRaw(requestParameters);
        return await response.value();
    }

    /**
     * User Endpoint for registration
     */
    async registerRaw(requestParameters: RegisterOperationRequest): Promise<runtime.ApiResponse<RegisterResponse>> {
        if (requestParameters.registerRequest === null || requestParameters.registerRequest === undefined) {
            throw new runtime.RequiredError('registerRequest','Required parameter requestParameters.registerRequest was null or undefined when calling register.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/users`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: RegisterRequestToJSON(requestParameters.registerRequest),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => RegisterResponseFromJSON(jsonValue));
    }

    /**
     * User Endpoint for registration
     */
    async register(requestParameters: RegisterOperationRequest): Promise<RegisterResponse> {
        const response = await this.registerRaw(requestParameters);
        return await response.value();
    }

}
