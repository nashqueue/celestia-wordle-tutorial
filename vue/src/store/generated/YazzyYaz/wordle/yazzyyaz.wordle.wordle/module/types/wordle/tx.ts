/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";

export const protobufPackage = "yazzyyaz.wordle.wordle";

export interface MsgSubmitWordle {
  creator: string;
  word: string;
}

export interface MsgSubmitWordleResponse {}

const baseMsgSubmitWordle: object = { creator: "", word: "" };

export const MsgSubmitWordle = {
  encode(message: MsgSubmitWordle, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.word !== "") {
      writer.uint32(18).string(message.word);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSubmitWordle {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgSubmitWordle } as MsgSubmitWordle;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.word = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSubmitWordle {
    const message = { ...baseMsgSubmitWordle } as MsgSubmitWordle;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.word !== undefined && object.word !== null) {
      message.word = String(object.word);
    } else {
      message.word = "";
    }
    return message;
  },

  toJSON(message: MsgSubmitWordle): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.word !== undefined && (obj.word = message.word);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgSubmitWordle>): MsgSubmitWordle {
    const message = { ...baseMsgSubmitWordle } as MsgSubmitWordle;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.word !== undefined && object.word !== null) {
      message.word = object.word;
    } else {
      message.word = "";
    }
    return message;
  },
};

const baseMsgSubmitWordleResponse: object = {};

export const MsgSubmitWordleResponse = {
  encode(_: MsgSubmitWordleResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSubmitWordleResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgSubmitWordleResponse,
    } as MsgSubmitWordleResponse;
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

  fromJSON(_: any): MsgSubmitWordleResponse {
    const message = {
      ...baseMsgSubmitWordleResponse,
    } as MsgSubmitWordleResponse;
    return message;
  },

  toJSON(_: MsgSubmitWordleResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgSubmitWordleResponse>
  ): MsgSubmitWordleResponse {
    const message = {
      ...baseMsgSubmitWordleResponse,
    } as MsgSubmitWordleResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  /** this line is used by starport scaffolding # proto/tx/rpc */
  SubmitWordle(request: MsgSubmitWordle): Promise<MsgSubmitWordleResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  SubmitWordle(request: MsgSubmitWordle): Promise<MsgSubmitWordleResponse> {
    const data = MsgSubmitWordle.encode(request).finish();
    const promise = this.rpc.request(
      "yazzyyaz.wordle.wordle.Msg",
      "SubmitWordle",
      data
    );
    return promise.then((data) =>
      MsgSubmitWordleResponse.decode(new Reader(data))
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
