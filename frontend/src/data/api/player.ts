import { rest } from '../rest';
import { PlayerResponse } from '../type';

const baseUri = '/player';

const getPlayer = async (playerID: string): Promise<PlayerResponse> => {
  try {
    const { data } = await rest.get<PlayerResponse>(`${baseUri}/${playerID}`);
    return data;
  } catch (error: any) {
    throw new Error(error);
  }
};

export { getPlayer };
