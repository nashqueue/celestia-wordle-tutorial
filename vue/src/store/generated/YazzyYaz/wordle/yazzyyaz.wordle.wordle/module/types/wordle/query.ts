/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../wordle/params";
import { Wordle } from "../wordle/wordle";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";

export const protobufPackage = "yazzyyaz.wordle.wordle";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetWordleRequest {
  index: string;
}

export interface QueryGetWordleResponse {
  wordle: Wordle | undefined;
}

export interface QueryAllWordleRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllWordleResponse {
  wordle: Wordle[];
  pagination: PageResponse | undefined;
}

const baseQueryParamsRequest: object = {};

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },
};

const baseQueryParamsResponse: object = {};

export const QueryParamsResponse = {
  encode(
    message: QueryParamsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },
};

const baseQueryGetWordleRequest: object = { index: "" };

export const QueryGetWordleRequest = {
  encode(
    message: QueryGetWordleRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetWordleRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetWordleRequest } as QueryGetWordleRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetWordleRequest {
    const message = { ...baseQueryGetWordleRequest } as QueryGetWordleRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: QueryGetWordleRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetWordleRequest>
  ): QueryGetWordleRequest {
    const message = { ...baseQueryGetWordleRequest } as QueryGetWordleRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseQueryGetWordleResponse: object = {};

export const QueryGetWordleResponse = {
  encode(
    message: QueryGetWordleResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.wordle !== undefined) {
      Wordle.encode(message.wordle, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetWordleResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetWordleResponse } as QueryGetWordleResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.wordle = Wordle.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetWordleResponse {
    const message = { ...baseQueryGetWordleResponse } as QueryGetWordleResponse;
    if (object.wordle !== undefined && object.wordle !== null) {
      message.wordle = Wordle.fromJSON(object.wordle);
    } else {
      message.wordle = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetWordleResponse): unknown {
    const obj: any = {};
    message.wordle !== undefined &&
      (obj.wordle = message.wordle ? Wordle.toJSON(message.wordle) : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetWordleResponse>
  ): QueryGetWordleResponse {
    const message = { ...baseQueryGetWordleResponse } as QueryGetWordleResponse;
    if (object.wordle !== undefined && object.wordle !== null) {
      message.wordle = Wordle.fromPartial(object.wordle);
    } else {
      message.wordle = undefined;
    }
    return message;
  },
};

const baseQueryAllWordleRequest: object = {};

export const QueryAllWordleRequest = {
  encode(
    message: QueryAllWordleRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllWordleRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllWordleRequest } as QueryAllWordleRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllWordleRequest {
    const message = { ...baseQueryAllWordleRequest } as QueryAllWordleRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllWordleRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllWordleRequest>
  ): QueryAllWordleRequest {
    const message = { ...baseQueryAllWordleRequest } as QueryAllWordleRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllWordleResponse: object = {};

export const QueryAllWordleResponse = {
  encode(
    message: QueryAllWordleResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.wordle) {
      Wordle.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllWordleResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllWordleResponse } as QueryAllWordleResponse;
    message.wordle = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.wordle.push(Wordle.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllWordleResponse {
    const message = { ...baseQueryAllWordleResponse } as QueryAllWordleResponse;
    message.wordle = [];
    if (object.wordle !== undefined && object.wordle !== null) {
      for (const e of object.wordle) {
        message.wordle.push(Wordle.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllWordleResponse): unknown {
    const obj: any = {};
    if (message.wordle) {
      obj.wordle = message.wordle.map((e) =>
        e ? Wordle.toJSON(e) : undefined
      );
    } else {
      obj.wordle = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllWordleResponse>
  ): QueryAllWordleResponse {
    const message = { ...baseQueryAllWordleResponse } as QueryAllWordleResponse;
    message.wordle = [];
    if (object.wordle !== undefined && object.wordle !== null) {
      for (const e of object.wordle) {
        message.wordle.push(Wordle.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a Wordle by index. */
  Wordle(request: QueryGetWordleRequest): Promise<QueryGetWordleResponse>;
  /** Queries a list of Wordle items. */
  WordleAll(request: QueryAllWordleRequest): Promise<QueryAllWordleResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "yazzyyaz.wordle.wordle.Query",
      "Params",
      data
    );
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
  }

  Wordle(request: QueryGetWordleRequest): Promise<QueryGetWordleResponse> {
    const data = QueryGetWordleRequest.encode(request).finish();
    const promise = this.rpc.request(
      "yazzyyaz.wordle.wordle.Query",
      "Wordle",
      data
    );
    return promise.then((data) =>
      QueryGetWordleResponse.decode(new Reader(data))
    );
  }

  WordleAll(request: QueryAllWordleRequest): Promise<QueryAllWordleResponse> {
    const data = QueryAllWordleRequest.encode(request).finish();
    const promise = this.rpc.request(
      "yazzyyaz.wordle.wordle.Query",
      "WordleAll",
      data
    );
    return promise.then((data) =>
      QueryAllWordleResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;
