import 'package:workspace/app/modules/cards/domain/exception/card_exception.dart';
import 'package:workspace/app/modules/cards/infrastructure/datasource/i_card_datasource.dart';
import 'package:workspace/app/shared/services/network/network_exception.dart';
import 'package:workspace/app/shared/services/network/network_service.dart';

class CardDatasourceImpl implements ICardDatasource {
  final NetworkService _networkService;

  CardDatasourceImpl(this._networkService);
  @override
  Future<bool> Create(Map<String, dynamic> json) async {
    try {
      await _networkService.post("/card/create", data: json);
      return true;
    } on NetworkException catch (e, s) {
      throw CardException(e.logger.toString(), s);
    } catch (e, s) {
      throw CardException("Error: 500 Internal Server Error", s);
    }
  }

  @override
  Future<bool> Delete(String id) async {
    try {
      await _networkService.delete("/card/delete/${id}");
      return true;
    } on NetworkException catch (e, s) {
      throw CardException(e.logger.toString(), s);
    } catch (e, s) {
      throw CardException("Error: 500 Internal Server Error", s);
    }
  }

  @override
  Future<dynamic> Find(int limit, int amount, String? search) async {
    try {
      final result = await _networkService
          .get("/card/list?limit=${limit}&offset=${amount}&search=${search}");
      return result.data["data"];
    } on NetworkException catch (e, s) {
      throw CardException(e.logger.toString(), s);
    } catch (e, s) {
      throw CardException("Error: 500 Internal Server Error", s);
    }
  }

  @override
  Future<bool> Update(Map<String, dynamic> json) async {
    try {
      await _networkService.patch("/card/update", data: json);
      return true;
    } on NetworkException catch (e, s) {
      throw CardException(e.logger.toString(), s);
    } catch (e, s) {
      throw CardException("Error: 500 Internal Server Error", s);
    }
  }
}
