import 'dart:ffi';

import 'package:app/generated/protobuf/content/content.enum.pb.dart';
import 'package:app/global/controller.dart';
import 'package:app/pages/comment/comment_add_controller.dart';
import 'package:app/pages/home/home_binding.dart';
import 'package:app/pages/home/home_view.dart';
import 'package:app/pages/home/splash_view.dart';
import 'package:app/pages/moment/detail/moment_detail_view.dart';
import 'package:app/pages/user/login_view.dart';
import 'package:app/pages/moment/add/moment_add_controller.dart';
import 'package:app/pages/moment/add/moment_add_view.dart';
import 'package:app/pages/moment/list/moment_list_view.dart';
import 'package:app/pages/moment/moment_binding.dart';
import 'package:app/pages/webview/webview.dart';
import 'package:get/get.dart';
import 'package:fixnum/fixnum.dart';

class AppPages {
  static final routes = [
    GetPage(
      name: Routes.HOME,
      page: () => HomeView(),
      bindings: [HomeBinding(),MomentBinding()],
    ),
    GetPage(
      name: Routes.MOMENT,
      page: () => MomentListView(),
      binding: MomentBinding(),
      children:[
        GetPage(
          name: Routes.ADD,
          page: () =>  globalState.authCheck() ?? MomentAddView(),
          binding: BindingsBuilder.put(() => MediaAddController())
        ),
        GetPage(
            name: Routes.DynamicId,
            page: () =>  MomentDetailView(),
            binding: BindingsBuilder.put(() => MediaAddController())
        ),
      ]
    ),
    GetPage(
      name: Routes.LOGIN,
      page: () => LoginView(),
    ),
    GetPage(
      name: Routes.SPLASH,
      page: () => Splash(),
    ),
    GetPage(
      name: Routes.WEBVIEW,
      page: () => WebViewExample(),
    ),
  ];
}

abstract class Routes {
  Routes._();

  static const HOME = '/home';
  static const CONTENT = '/content';
  static const MOMENT = '/moment';
  static const ADD = '/add';
  static const MOMENT_ADD = MOMENT + ADD;
  static const LOGIN = '/login';
  static const SETTINGS = '/settings';
  static const SPLASH = '/splash';
  static const PRODUCT = '/product';
  static const DynamicId = '/:id';
  static const PRODUCT_DETAILS = PRODUCT + DynamicId;
  static const MOMENT_DETAILS = MOMENT + DynamicId;
  static const WEBVIEW = '/webview';
  static const NOTFOUND = '/NOTFOUND';

  static String productDetails(String productId) => '$PRODUCT/$productId';
  static String momentDetails(String momentId) => '$MOMENT/$momentId';
  static String contentDetails(ContentType type,Int64 contentId) => '${getContentRoute(type)}/$contentId';

  static  String getContentRoute(ContentType type) {
    switch (type) {
      case ContentType.ContentMoment:return Routes.MOMENT;
    }
    return Routes.NOTFOUND;
  }
}

