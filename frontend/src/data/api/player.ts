import { rest } from '../rest';
import { PlayerResponse, PlayersResponse } from '../type';

const baseUri = '/player';

const getPlayer = async (playerID: string): Promise<PlayerResponse> => {
  try {
    const { data } = await rest.get<PlayerResponse>(`${baseUri}/${playerID}`);
    return data;
  } catch (error: any) {
    throw new Error(error);
  }
};

const searchPlayer = async (name: string): Promise<PlayersResponse> => {
  try {
    const { data } = await rest.getParams(baseUri + '/search', {
      params: {
        Name: name,
      },
    });
    return data;
  } catch (error: any) {
    throw new Error(error);
  }
};

export { getPlayer, searchPlayer };
