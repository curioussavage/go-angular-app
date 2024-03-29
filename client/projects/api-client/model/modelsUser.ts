/**
 * User API
 * This is an API for managing users.
 *
 * OpenAPI spec version: 1.0
 * Contact: support@spiderman.com
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 * Do not edit the class manually.
 */
import { ModelsUserStatus } from './modelsUserStatus';

/**
 * User account information
 */
export interface ModelsUser { 
    department?: string;
    email?: string;
    firstName?: string;
    lastName?: string;
    userID?: number;
    userName?: string;
    userStatus?: ModelsUserStatus;
}