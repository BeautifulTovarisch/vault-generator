'use strict';

import { post } from 'axios';

export const encryptConfig = payload => post("/v0/api/vault", payload);
