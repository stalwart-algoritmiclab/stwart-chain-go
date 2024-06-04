/* eslint-disable */
/* tslint:disable */
/*
 * ---------------------------------------------------------------
 * ## THIS FILE WAS GENERATED VIA SWAGGER-TYPESCRIPT-API        ##
 * ##                                                           ##
 * ## AUTHOR: acacode                                           ##
 * ## SOURCE: https://github.com/acacode/swagger-typescript-api ##
 * ---------------------------------------------------------------
 */

export interface Any {
  "@type"?: string;
}

export interface Status {
  /** @format int32 */
  code?: number;
  message?: string;
  details?: { "@type"?: string }[];
}

export interface Address {
  /** @format uint64 */
  id?: string;
  address?: string;
  creator?: string;
}

export interface Fees {
  amountFrom?: string;
  fee?: string;
  refReward?: string;
  stakeReward?: string;

  /** @format uint64 */
  minAmount?: string;
  noRefReward?: boolean;
  creator?: string;

  /** @format uint64 */
  id?: string;
}

export interface PageRequest {
  /** @format byte */
  key?: string;

  /** @format uint64 */
  offset?: string;

  /** @format uint64 */
  limit?: string;
  count_total?: boolean;
  reverse?: boolean;
}

export interface PageResponse {
  /** @format byte */
  next_key?: string;

  /** @format uint64 */
  total?: string;
}

export interface QueryAllAddressesResponse {
  Addresses?: { id?: string; address?: string; creator?: string }[];
  pagination?: { next_key?: string; total?: string };
}

export interface QueryAllTariffResponse {
  tariff?: {
    denom?: string;
    id?: string;
    amount?: string;
    minRefBalance?: string;
    fees?: {
      amountFrom?: string;
      fee?: string;
      refReward?: string;
      stakeReward?: string;
      minAmount?: string;
      noRefReward?: boolean;
      creator?: string;
      id?: string;
    }[];
  }[];
  pagination?: { next_key?: string; total?: string };
}

export interface QueryAllTariffsResponse {
  tariffs?: {
    denom?: string;
    tariffs?: {
      denom?: string;
      id?: string;
      amount?: string;
      minRefBalance?: string;
      fees?: {
        amountFrom?: string;
        fee?: string;
        refReward?: string;
        stakeReward?: string;
        minAmount?: string;
        noRefReward?: boolean;
        creator?: string;
        id?: string;
      }[];
    }[];
    creator?: string;
  }[];
  pagination?: { next_key?: string; total?: string };
}

export interface QueryGetAddressesResponse {
  Addresses?: { id?: string; address?: string; creator?: string };
}

export interface QueryGetTariffResponse {
  tariff?: {
    denom?: string;
    id?: string;
    amount?: string;
    minRefBalance?: string;
    fees?: {
      amountFrom?: string;
      fee?: string;
      refReward?: string;
      stakeReward?: string;
      minAmount?: string;
      noRefReward?: boolean;
      creator?: string;
      id?: string;
    }[];
  };
}

export interface QueryGetTariffsResponse {
  tariffs?: {
    denom?: string;
    tariffs?: {
      denom?: string;
      id?: string;
      amount?: string;
      minRefBalance?: string;
      fees?: {
        amountFrom?: string;
        fee?: string;
        refReward?: string;
        stakeReward?: string;
        minAmount?: string;
        noRefReward?: boolean;
        creator?: string;
        id?: string;
      }[];
    }[];
    creator?: string;
  };
}

export interface QueryParamsResponse {
  params?: object;
}

export type FeepolicyParams = object;

export interface FeepolicyTariff {
  denom?: string;

  /** @format uint64 */
  id?: string;
  amount?: string;
  minRefBalance?: string;
  fees?: {
    amountFrom?: string;
    fee?: string;
    refReward?: string;
    stakeReward?: string;
    minAmount?: string;
    noRefReward?: boolean;
    creator?: string;
    id?: string;
  }[];
}

export interface FeepolicyTariffs {
  denom?: string;
  tariffs?: {
    denom?: string;
    id?: string;
    amount?: string;
    minRefBalance?: string;
    fees?: {
      amountFrom?: string;
      fee?: string;
      refReward?: string;
      stakeReward?: string;
      minAmount?: string;
      noRefReward?: boolean;
      creator?: string;
      id?: string;
    }[];
  }[];
  creator?: string;
}

export interface MsgCreateAddressesResponse {
  /** @format uint64 */
  id?: string;
}

export type MsgCreateTariffsResponse = object;

export type MsgDeleteAddressesResponse = object;

export type MsgDeleteTariffsResponse = object;

export type MsgUpdateAddressesResponse = object;

export type MsgUpdateParamsResponse = object;

export type MsgUpdateTariffsResponse = object;

export type Params = object;

export interface Tariff {
  denom?: string;

  /** @format uint64 */
  id?: string;
  amount?: string;
  minRefBalance?: string;
  fees?: {
    amountFrom?: string;
    fee?: string;
    refReward?: string;
    stakeReward?: string;
    minAmount?: string;
    noRefReward?: boolean;
    creator?: string;
    id?: string;
  }[];
}

import axios, { AxiosInstance, AxiosRequestConfig, AxiosResponse, ResponseType } from "axios";

export type QueryParamsType = Record<string | number, any>;

export interface FullRequestParams extends Omit<AxiosRequestConfig, "data" | "params" | "url" | "responseType"> {
  /** set parameter to `true` for call `securityWorker` for this request */
  secure?: boolean;
  /** request path */
  path: string;
  /** content type of request body */
  type?: ContentType;
  /** query params */
  query?: QueryParamsType;
  /** format of response (i.e. response.json() -> format: "json") */
  format?: ResponseType;
  /** request body */
  body?: unknown;
}

export type RequestParams = Omit<FullRequestParams, "body" | "method" | "query" | "path">;

export interface ApiConfig<SecurityDataType = unknown> extends Omit<AxiosRequestConfig, "data" | "cancelToken"> {
  securityWorker?: (
    securityData: SecurityDataType | null,
  ) => Promise<AxiosRequestConfig | void> | AxiosRequestConfig | void;
  secure?: boolean;
  format?: ResponseType;
}

export enum ContentType {
  Json = "application/json",
  FormData = "multipart/form-data",
  UrlEncoded = "application/x-www-form-urlencoded",
}

export class HttpClient<SecurityDataType = unknown> {
  public instance: AxiosInstance;
  private securityData: SecurityDataType | null = null;
  private securityWorker?: ApiConfig<SecurityDataType>["securityWorker"];
  private secure?: boolean;
  private format?: ResponseType;

  constructor({ securityWorker, secure, format, ...axiosConfig }: ApiConfig<SecurityDataType> = {}) {
    this.instance = axios.create({ ...axiosConfig, baseURL: axiosConfig.baseURL || "" });
    this.secure = secure;
    this.format = format;
    this.securityWorker = securityWorker;
  }

  public setSecurityData = (data: SecurityDataType | null) => {
    this.securityData = data;
  };

  private mergeRequestParams(params1: AxiosRequestConfig, params2?: AxiosRequestConfig): AxiosRequestConfig {
    return {
      ...this.instance.defaults,
      ...params1,
      ...(params2 || {}),
      headers: {
        ...(this.instance.defaults.headers || {}),
        ...(params1.headers || {}),
        ...((params2 && params2.headers) || {}),
      },
    };
  }

  private createFormData(input: Record<string, unknown>): FormData {
    return Object.keys(input || {}).reduce((formData, key) => {
      const property = input[key];
      formData.append(
        key,
        property instanceof Blob
          ? property
          : typeof property === "object" && property !== null
          ? JSON.stringify(property)
          : `${property}`,
      );
      return formData;
    }, new FormData());
  }

  public request = async <T = any, _E = any>({
    secure,
    path,
    type,
    query,
    format,
    body,
    ...params
  }: FullRequestParams): Promise<AxiosResponse<T>> => {
    const secureParams =
      ((typeof secure === "boolean" ? secure : this.secure) &&
        this.securityWorker &&
        (await this.securityWorker(this.securityData))) ||
      {};
    const requestParams = this.mergeRequestParams(params, secureParams);
    const responseFormat = (format && this.format) || void 0;

    if (type === ContentType.FormData && body && body !== null && typeof body === "object") {
      requestParams.headers.common = { Accept: "*/*" };
      requestParams.headers.post = {};
      requestParams.headers.put = {};

      body = this.createFormData(body as Record<string, unknown>);
    }

    return this.instance.request({
      ...requestParams,
      headers: {
        ...(type && type !== ContentType.FormData ? { "Content-Type": type } : {}),
        ...(requestParams.headers || {}),
      },
      params: query,
      responseType: responseFormat,
      data: body,
      url: path,
    });
  };
}

/**
 * @title HTTP API Console stwartchain.feepolicy
 */
export class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
  /**
   * No description
   *
   * @tags Query
   * @name QueryAddressesAll
   * @request GET:/stwart/feepolicy/addresses
   */
  queryAddressesAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<
      {
        Addresses?: { id?: string; address?: string; creator?: string }[];
        pagination?: { next_key?: string; total?: string };
      },
      { code?: number; message?: string; details?: { "@type"?: string }[] }
    >({
      path: `/stwart/feepolicy/addresses`,
      method: "GET",
      query: query,
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryAddresses
   * @request GET:/stwart/feepolicy/addresses/{id}
   */
  queryAddresses = (id: string, params: RequestParams = {}) =>
    this.request<
      { Addresses?: { id?: string; address?: string; creator?: string } },
      { code?: number; message?: string; details?: { "@type"?: string }[] }
    >({
      path: `/stwart/feepolicy/addresses/${id}`,
      method: "GET",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryParams
   * @request GET:/stwart/feepolicy/params
   */
  queryParams = (params: RequestParams = {}) =>
    this.request<{ params?: object }, { code?: number; message?: string; details?: { "@type"?: string }[] }>({
      path: `/stwart/feepolicy/params`,
      method: "GET",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryTariffAll
   * @request GET:/stwart/feepolicy/tariff
   */
  queryTariffAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<
      {
        tariff?: {
          denom?: string;
          id?: string;
          amount?: string;
          minRefBalance?: string;
          fees?: {
            amountFrom?: string;
            fee?: string;
            refReward?: string;
            stakeReward?: string;
            minAmount?: string;
            noRefReward?: boolean;
            creator?: string;
            id?: string;
          }[];
        }[];
        pagination?: { next_key?: string; total?: string };
      },
      { code?: number; message?: string; details?: { "@type"?: string }[] }
    >({
      path: `/stwart/feepolicy/tariff`,
      method: "GET",
      query: query,
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryTariff
   * @request GET:/stwart/feepolicy/tariff/{denom}
   */
  queryTariff = (denom: string, params: RequestParams = {}) =>
    this.request<
      {
        tariff?: {
          denom?: string;
          id?: string;
          amount?: string;
          minRefBalance?: string;
          fees?: {
            amountFrom?: string;
            fee?: string;
            refReward?: string;
            stakeReward?: string;
            minAmount?: string;
            noRefReward?: boolean;
            creator?: string;
            id?: string;
          }[];
        };
      },
      { code?: number; message?: string; details?: { "@type"?: string }[] }
    >({
      path: `/stwart/feepolicy/tariff/${denom}`,
      method: "GET",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryTariffsAll
   * @request GET:/stwart/feepolicy/tariffs
   */
  queryTariffsAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<
      {
        tariffs?: {
          denom?: string;
          tariffs?: {
            denom?: string;
            id?: string;
            amount?: string;
            minRefBalance?: string;
            fees?: {
              amountFrom?: string;
              fee?: string;
              refReward?: string;
              stakeReward?: string;
              minAmount?: string;
              noRefReward?: boolean;
              creator?: string;
              id?: string;
            }[];
          }[];
          creator?: string;
        }[];
        pagination?: { next_key?: string; total?: string };
      },
      { code?: number; message?: string; details?: { "@type"?: string }[] }
    >({
      path: `/stwart/feepolicy/tariffs`,
      method: "GET",
      query: query,
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryTariffs
   * @request GET:/stwart/feepolicy/tariffs/{denom}
   */
  queryTariffs = (denom: string, params: RequestParams = {}) =>
    this.request<
      {
        tariffs?: {
          denom?: string;
          tariffs?: {
            denom?: string;
            id?: string;
            amount?: string;
            minRefBalance?: string;
            fees?: {
              amountFrom?: string;
              fee?: string;
              refReward?: string;
              stakeReward?: string;
              minAmount?: string;
              noRefReward?: boolean;
              creator?: string;
              id?: string;
            }[];
          }[];
          creator?: string;
        };
      },
      { code?: number; message?: string; details?: { "@type"?: string }[] }
    >({
      path: `/stwart/feepolicy/tariffs/${denom}`,
      method: "GET",
      ...params,
    });
}
