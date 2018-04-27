#include <eosiolib/eosio.hpp>
#include <eosiolib/print.hpp>

using namespace eosio;
using namespace std;

class configure : public eosio::contract {
  public:
      using contract::contract;

      /// @abi action 
      void hi( account_name user ) {
         print( "Hello, ", name{user} );
         config conf;
         get_config(conf);
         print("conf.is_manual_setprods: ", conf.is_manual_setprods);
         if (conf.is_manual_setprods == 0) {
          conf.is_manual_setprods = 1;
        } else {
          conf.is_manual_setprods = 0;
        }
        store_config(conf);
      }

  private:
    struct config {
      config() {}
      constexpr static uint64_t key = N(config);
      uint64_t is_manual_setprods = 0;
    };

    void store_config(const config &conf) {
      auto it = db_find_i64(_self, _self, N(config), config::key);
      if (it != -1) {
        db_update_i64(it, _self, (const char *)&conf, sizeof(config));
      } else {
        db_store_i64(_self, N(config), _self, config::key, (const char *)&conf, sizeof(config));
      }
    }

    bool get_config(config &conf) {
      auto it = db_find_i64(_self, _self, N(config), config::key);
      if (it != -1) {
        auto size = db_get_i64(it, (char*)&conf, sizeof(config));
        eosio_assert(size == sizeof(config), "Wrong record size");
        return true;
      }
      return false;
    }
};

EOSIO_ABI(configure, (hi))
